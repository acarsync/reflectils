package reflectils_test

import (
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

func ExampleGetValueByKind() {
	v := reflect.ValueOf(&ExampleStruct{
		A: "reflectils",
	})
	res, ok := reflectils.GetValueByKind(v, reflect.Struct)
	fmt.Println(ok, res)
	// Output:
	// true {reflectils 0 0}
}

func ExampleGetValueByKind_withInvalidValue() {
	res, ok := reflectils.GetValueByKind(reflect.ValueOf("reflectils"), reflect.Bool)
	fmt.Println(ok, res)
	// Output:
	// false <invalid reflect.Value>
}
