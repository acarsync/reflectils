package reflectils_test

import (
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

func ExampleGetValueByKind() {
	v := reflect.ValueOf(&struct{ Name string }{"reflectils"})
	res, ok := reflectils.GetValueByKind(v, reflect.Struct)
	fmt.Println(v)
	fmt.Println(ok, res)
	// Output:
	// &{reflectils}
	// true {reflectils}
}

func ExampleGetValueByKind_withInvalidValue() {
	res, ok := reflectils.GetValueByKind(reflect.ValueOf("reflectils"), reflect.Bool)
	fmt.Println(ok, res)
	// Output:
	// false <invalid reflect.Value>
}
