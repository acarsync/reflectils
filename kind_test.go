package reflectils_test

import (
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

func ExampleMapOfKind() {
	m := reflectils.NewMapOfKind(reflect.String, reflect.Int)
	fmt.Println(m.Has(reflect.String))
	fmt.Println(m.Has(reflect.Uint))
	fmt.Println(m.Has(reflect.Int))
	// Output:
	// true
	// false
	// true
}

func ExampleAnyKind() {
	v := reflect.ValueOf("reflectils")
	fmt.Println(reflectils.AnyKind(v, reflect.String))
	fmt.Println(reflectils.AnyKind(v, reflect.Uint))
	// Output:
	// true
	// false
}
