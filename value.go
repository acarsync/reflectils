package reflectils

import (
	"errors"
	"reflect"
)

// ErrStopIteration tells IterValue to break iteration.
var ErrStopIteration = errors.New("ErrStopIteration")

// GetValueByKind returns v if v type matchs k.
// will use v.Elem if v type is ptr.
func GetValueByKind(v reflect.Value, k reflect.Kind) (reflect.Value, bool) {
again:
	switch v.Kind() {
	case k:
		return v, true
	case reflect.Ptr:
		v = v.Elem()
		goto again
	}
	return reflect.Value{}, false
}

// NewSliceOfValue returns new slice of reflect.Value from v.
func NewSliceOfValue(v ...interface{}) (r []reflect.Value) {
	for _, e := range v {
		if ve, ok := e.(reflect.Value); ok {
			r = append(r, ve)
		} else {
			r = append(r, reflect.ValueOf(e))
		}
	}
	return
}

// IterValue iters v until f returns an error.
// Return err will be nil if f returns ErrStopIteration.
func IterValue(v []reflect.Value, f func(i int, e reflect.Value) (err error)) (err error) {
	for i, e := range v {
		if err = f(i, e); err != nil {
			break
		}
	}
	if errors.Is(err, ErrStopIteration) {
		err = nil
	}
	return
}
