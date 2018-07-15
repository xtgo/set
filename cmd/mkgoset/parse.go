package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func Plan(importPath, dst, src string) (*PlanData, error) {
	fset := token.NewFileSet()
	pkg, err := parsePackage(fset, importPath)
	if err != nil {
		return nil, err
	}

	p := &PlanData{SetType: dst, PkgScope: pkg.Scope()}
	err = p.decode(src)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func localQualifier(p *types.Package) string { return "" }

func parsePackage(fset *token.FileSet, dir string) (*types.Package, error) {
	bpkg, err := build.Default.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}

	files := make([]*ast.File, len(bpkg.GoFiles))
	for i, path := range bpkg.GoFiles {
		files[i], err = parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil, err
		}
	}
	conf := types.Config{
		IgnoreFuncBodies:         true,
		DisableUnusedImportCheck: true,
		Importer:                 importer.Default(),
	}
	return conf.Check(bpkg.Dir, fset, files, nil)
}

type PlanData struct {
	SetType, ElemType, FuncName, FuncDef string

	PkgScope *types.Scope
}

func (p *PlanData) decode(sym string) error {
	_, obj := p.PkgScope.LookupParent(sym, token.NoPos)
	switch v := obj.(type) {
	case nil:
		return fmt.Errorf("symbol %#q not found in package", sym)
	default:
		return fmt.Errorf("%#q is not one of ElemType, LessFn, or SelFn", sym)
	case *types.TypeName:
		return p.decodeElemType(v)
	case *types.Func:
		return p.decodeFunc(v)
	}
}

func (p *PlanData) decodeElemType(tn *types.TypeName) error {
	sym := tn.Name()
	if !isOrdered(tn.Type()) {
		return fmt.Errorf("ElemType %#q is not compatible with the '<' operator", sym)
	}
	funcName := "less" + p.SetType
	format := `func %s(a, b %s) bool { return a < b }`
	p.ElemType = sym
	p.FuncName = funcName
	p.FuncDef = fmt.Sprintf(format, funcName, sym)
	return nil
}

func (p *PlanData) decodeFunc(f *types.Func) error {
	sig := funcSig(f)
	switch {
	case isLessFn(sig):
		return p.decodeLessFn(f)
	case isSelFn(sig):
		return p.decodeSelFn(f)
	default:
		return fmt.Errorf("%#q is neither a `func(T, T) bool` nor `func(T1) T2`", funcRepr(f))
	}
}

func (p *PlanData) decodeLessFn(f *types.Func) error {
	err := checkLessFn(f)
	if err != nil {
		return err
	}
	sig := funcSig(f)
	p.ElemType = typeRepr(sig.Params().At(0).Type())
	p.FuncName = f.Name()
	return nil
}

func (p *PlanData) decodeSelFn(f *types.Func) error {
	err := checkSelFn(f)
	if err != nil {
		return err
	}
	sig := funcSig(f)
	outer := typeRepr(sig.Params().At(0).Type())
	inner := typeRepr(sig.Results().At(0).Type())
	funcName := "less" + p.SetType
	format := `func %s(a, b %s) bool { return %[3]s(a) < %[3]s(b) }`
	p.ElemType = inner
	p.FuncName = funcName
	p.FuncDef = fmt.Sprintf(format, funcName, outer, f.Name())
	return nil
}

func funcRepr(f *types.Func) string { return f.Name() }
func typeRepr(t types.Type) string  { return types.TypeString(t, localQualifier) }

func arity(sig *types.Signature, nparams, nresults int) bool {
	return sig.Params().Len() == nparams && sig.Results().Len() == nresults
}

func isOrdered(t types.Type) bool { return hasInfo(t, types.IsOrdered) }
func isBoolean(t types.Type) bool { return hasInfo(t, types.IsBoolean) }

func isLessFn(sig *types.Signature) bool { return arity(sig, 2, 1) }
func isSelFn(sig *types.Signature) bool  { return arity(sig, 1, 1) }

func hasInfo(t types.Type, info types.BasicInfo) bool {
	basic, ok := t.Underlying().(*types.Basic)
	return ok && basic.Info()&info != 0
}

func funcSig(f *types.Func) *types.Signature {
	return f.Type().(*types.Signature)
}

func checkLessFn(f *types.Func) error {
	sig := funcSig(f)
	if !isLessFn(sig) {
		return fmt.Errorf("%#q is not a `func(T, T) bool`", funcRepr(f))
	}
	ps, rs := sig.Params(), sig.Results()
	p0, p1, r0 := ps.At(0).Type(), ps.At(1).Type(), rs.At(0).Type()
	switch {
	case p0 != p1:
		return fmt.Errorf("%#q inputs must have same type (%#q != %#q)",
			funcRepr(f), typeRepr(p0), typeRepr(p1))
	case !isBoolean(r0):
		return fmt.Errorf("%#q must return bool (not %#q)",
			funcRepr(f), typeRepr(r0))
	}
	return nil
}

func checkSelFn(f *types.Func) error {
	sig := funcSig(f)
	if !isSelFn(sig) {
		return fmt.Errorf("%#q is not a `func(T1) T2`", f)
	}
	r0 := sig.Results().At(0).Type()
	if !isOrdered(r0) {
		return fmt.Errorf("SelFn return type %#q is not compatible with the '<' operator", r0)
	}
	return nil
}
