package reflectils

import (
	"reflect"
	"unsafe"
)

// StringFromBytes returns string from b without memory allocation.
func StringFromBytes(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes returns bytes from s without memory allocation.
func StringToBytes(s string) []byte {
	b := *(*[]byte)(unsafe.Pointer(&s))
	(*reflect.SliceHeader)(unsafe.Pointer(&b)).Cap = len(s)
	return b
}
