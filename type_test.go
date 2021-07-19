package reflectils_test

import (
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

func ExampleGetTypeByKind() {
	v := reflect.TypeOf(&struct{ Name string }{"reflectils"})
	res, ok := reflectils.GetTypeByKind(v, reflect.Struct)
	fmt.Println(v)
	fmt.Println(ok, res)
	// Output:
	// *struct { Name string }
	// true struct { Name string }
}

func ExampleGetTypeByKind_withInvalidValue() {
	res, ok := reflectils.GetTypeByKind(reflect.TypeOf("reflectils"), reflect.Bool)
	fmt.Println(ok, res)
	// Output:
	// false <nil>
}
