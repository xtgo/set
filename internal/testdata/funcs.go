// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import "reflect"

// IsEqual is a simple int slice equality implementation.
func IsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ConvField converts and stores the value of src.field into dst. dst must
// be a pointer.
func ConvField(dst interface{}, src interface{}, field string) {
	v1 := reflect.ValueOf(dst).Elem()
	v2 := reflect.ValueOf(src)
	v2 = v2.FieldByName(field)
	v2 = v2.Convert(v1.Type())
	v1.Set(v2)
}

// ConvMethod converts and stores the method expression of type(src).method
// into dst. dst must be a pointer to function type.
func ConvMethod(dst interface{}, src interface{}, method string) {
	v1 := reflect.ValueOf(dst).Elem()
	t := reflect.TypeOf(src)
	op, _ := t.MethodByName(method)
	v2 := op.Func.Convert(v1.Type())
	v1.Set(v2)
}
