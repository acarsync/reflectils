package reflectils

import (
	"reflect"
)

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
