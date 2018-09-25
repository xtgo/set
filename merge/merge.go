package merge

import "sort"

type Action byte

const (
	SkipL Action = 1 << iota // Advance Left
	SkipR                    // Advance Right
	KeepL
	KeepR
	maxAction

	keeps = KeepL | KeepR
)

func (a Action) IsValid() bool {
	return a != 0 && a < maxAction &&
		a&keeps != keeps // specifying both keeps is prohibited
}

type Table struct {
	// Default Action is "retry"
	LT Action // Left is Less-Than Right
	Eq Action // Left is Equal to Right
	GT Action // Left is Greater-Than Right

	// Only Left (use *L actions) or Right (use *R actions) remain.
	// Default Action is Skip*.
	Rem Action
}

func (t Table) IsValid() bool {
	// Rem is always valid
	return t.LT.IsValid() && t.EQ.IsValid() && t.GT.IsValid() &&
		t.Rem == 0 || t.Rem.IsValid()
}

type lesser interface {
	Less(i, j int) bool
}

func (t Table) cmp(l lesser, i, j int) Action {
	switch {
	case l.Less(i, j):
		return t.LT
	case l.Less(j, i):
		return t.GT
	default:
		return t.Eq
	}
}

type slider interface {
	Slide(i, j, k int)
}

type swapper interface {
	Swap(i, j int)
}

func sliderFrom(sw swapper) slider {
	sl, ok := sw.(slider)
	if ok {
		return sl
	}
	return swapSlider{sw}
}

type swapSlider struct{ swapper }

func (s swapSlider) Slide(i, j, n int) {
	for k := 0; k < n; k++ {
		s.Swap(i+k, j+k)
	}
}

func Merge(t Table, data sort.Interface, pivot int) (size int) {
	if !t.IsValid() {
		panic("merge: table not valid")
	}
	k, l := pivot, data.Len()
	i, j := 0, k
	p := 0 // index of next spot to store
	for i < k && j < l {
		a := t.cmp(data, i, j)
		switch {
		case a&KeepL != 0 && p != i:
			// TODO: what if p > i ?
			data.Swap(p, i)
			fallthrough
		case a.KeepL != 0:
			// TODO: what if p > i ?
			p++
			fallthrough
		case a&SkipL != 0:
			i++
		}
		switch {
		case a&KeepR != 0:
			// TODO: what if p > i ?
			data.Swap(p, j)
			p++
			fallthrough
		case a&SkipR != 0:
			j++
		}
	}

	sldr := sliderFrom(sw)
	switch {
	case i < k && t.Rem&KeepL != 0:
		if p != i {
			// TODO: what if p > i ?
			sldr.Slide(p, i, k-i)
		}
		// TODO: what if p > i ?
		p = k
	case j < l && t.Rem&KeepR != 0:
		sldr.Slide(p, j, l-j)
		p += l - j
	}
	return p
}

func Diff(data sort.Interface, pivot int) (size int) {
	t := Table{LT: KeepL, Eq: SkipL | SkipR, GT: SkipR, Rem: KeepL}
	return Merge(t, data, pivot)
}

func Inter(data sort.Interface, pivot int) (size int) {
	t := Table{LT: SkipL, Eq: KeepL | SkipR, GT: SkipR}
	return Merge(t, data, pivot)
}

func Union(data sort.Interface, pivot int) (size int) {
	// TODO dupes?
	t := Table{LT: KeepL, Eq: KeepL | SkipR, GT: KeepR, Rem: KeepL | KeepR}
	return Merge(t, data, pivot)
}

func SymDiff(data sort.Interface, pivot int) (size int) {
	// TODO dupes?
	t := Table{LT: KeepL, Eq: SkipL | SkipR, GT: KeepR, Rem: KeepL | KeepR}
	return Merge(t, data, pivot)
}
