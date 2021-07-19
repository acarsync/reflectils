package reflectils

import (
	"reflect"
)

// EmptyStruct a struct that contains nothing.
type EmptyStruct struct{}

// MapTypeOfStructFields calls f with v fields.
func MapTypeOfStructFields(v reflect.Type, f func(i int, v reflect.Type)) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		f(i, v.Field(i).Type)
	}
}

// MapValueOfStructFields calls f with v fields.
func MapValueOfStructFields(v reflect.Value, f func(i int, v reflect.Value)) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		f(i, v.Field(i))
	}
}

// FindTypeOfStructFieldsByKind returns a field type of v if matches by k.
func FindTypeOfStructFieldsByKind(v reflect.Type, k ...reflect.Kind) (r map[int]reflect.Type) {
	r = make(map[int]reflect.Type)
	m := NewMapOfKind(k...)
	MapTypeOfStructFields(v, func(i int, v reflect.Type) {
		if m.Has(v.Kind()) {
			r[i] = v
		}
	})
	return
}

// FindValueOfStructFieldsByKind returns a field value of v if field type matches by k.
func FindValueOfStructFieldsByKind(v reflect.Value, k ...reflect.Kind) (r map[int]reflect.Value) {
	r = make(map[int]reflect.Value)
	m := NewMapOfKind(k...)
	MapValueOfStructFields(v, func(i int, v reflect.Value) {
		if m.Has(v.Kind()) {
			r[i] = v
		}
	})
	return
}
