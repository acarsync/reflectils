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
