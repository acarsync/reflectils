package reflectils

import (
	"reflect"
)

// EmptyStruct a struct that contains nothing.
type EmptyStruct struct{}

// MapStructFields calls f with v fields.
// will break iteration if f returns an error.
func MapStructFields(v reflect.Type, f func(i int, v reflect.StructField) error) (err error) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		if err = f(i, v.Field(i)); err != nil {
			break
		}
	}
	return
}

// MapTypeOfStructFields calls f with v fields.
// will break iteration if f returns an error.
func MapTypeOfStructFields(v reflect.Type, f func(i int, v reflect.Type) error) (err error) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		if err = f(i, v.Field(i).Type); err != nil {
			break
		}
	}
	return
}

// MapValueOfStructFields calls f with v fields.
// will break iteration if f returns an error.
func MapValueOfStructFields(v reflect.Value, f func(i int, v reflect.Value) error) (err error) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		if err = f(i, v.Field(i)); err != nil {
			break
		}
	}
	return
}

// FindTypeOfStructFieldsByKind returns a field type of v if matches by k.
func FindTypeOfStructFieldsByKind(v reflect.Type, k ...reflect.Kind) (r map[int]reflect.Type) {
	r = make(map[int]reflect.Type)
	m := NewMapOfKind(k...)
	MapTypeOfStructFields(v, func(i int, v reflect.Type) error {
		if m.Has(v.Kind()) {
			r[i] = v
		}
		return nil
	})
	return
}

// FindValueOfStructFieldsByKind returns a field value of v if field type matches by k.
func FindValueOfStructFieldsByKind(v reflect.Value, k ...reflect.Kind) (r map[int]reflect.Value) {
	r = make(map[int]reflect.Value)
	m := NewMapOfKind(k...)
	MapValueOfStructFields(v, func(i int, v reflect.Value) error {
		if m.Has(v.Kind()) {
			r[i] = v
		}
		return nil
	})
	return
}
