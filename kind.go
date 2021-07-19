package reflectils

import (
	"reflect"
)

// MapOfKind a map that contains reflect.Kind.
type MapOfKind map[reflect.Kind]EmptyStruct

// NewMapOfKind returns new MapOfKind.
func NewMapOfKind(k ...reflect.Kind) (v MapOfKind) {
	v = make(MapOfKind)
	for _, i := range k {
		v[i] = EmptyStruct{}
	}
	return
}

// Has returns true if k is in m.
func (m MapOfKind) Has(k reflect.Kind) bool {
	_, has := m[k]
	return has
}
