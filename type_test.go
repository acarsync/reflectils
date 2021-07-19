package reflectils_test

import (
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

func ExampleGetTypeByKind() {
	v := reflect.TypeOf(&ExampleStruct{
		A: "reflectils",
	})
	res, ok := reflectils.GetTypeByKind(v, reflect.Struct)
	fmt.Println(ok, res)
	// Output:
	// true reflectils_test.ExampleStruct
}

func ExampleGetTypeByKind_withInvalidValue() {
	res, ok := reflectils.GetTypeByKind(reflect.TypeOf("reflectils"), reflect.Bool)
	fmt.Println(ok, res)
	// Output:
	// false <nil>
}
