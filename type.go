package reflectils

import (
	"reflect"
)

// GetTypeByKind returns v if v matchs k.
// will use v.Elem if v is ptr.
func GetTypeByKind(v reflect.Type, k reflect.Kind) (reflect.Type, bool) {
again:
	switch v.Kind() {
	case k:
		return v, true
	case reflect.Ptr:
		v = v.Elem()
		goto again
	}
	return nil, false
}
