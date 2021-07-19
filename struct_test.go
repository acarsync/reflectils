package reflectils_test

import (
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

type ExampleStruct struct {
	A string
	B uint
	C int
}

func ExampleMapTypeOfStructFields() {
	v := reflect.TypeOf((*ExampleStruct)(nil)).Elem()
	reflectils.MapTypeOfStructFields(v, func(i int, v reflect.Type) {
		fmt.Println(i, v)
	})
	// Output:
	// 0 string
	// 1 uint
	// 2 int
}

func ExampleMapValueOfStructFields() {
	v := reflect.ValueOf(ExampleStruct{
		A: "reflectils",
		B: 123,
		C: 456,
	})
	reflectils.MapValueOfStructFields(v, func(i int, v reflect.Value) {
		fmt.Println(i, v)
	})
	// Output:
	// 0 reflectils
	// 1 123
	// 2 456
}

func ExampleFindTypeOfStructFieldsByKind() {
	v := reflect.TypeOf((*ExampleStruct)(nil)).Elem()
	fmt.Println(reflectils.FindTypeOfStructFieldsByKind(v, reflect.String, reflect.Int))
	// Output: map[0:string 2:int]
}

func ExampleFindValueOfStructFieldsByKind() {
	v := reflect.ValueOf(ExampleStruct{
		A: "reflectils",
		B: 123,
		C: 456,
	})
	fmt.Println(reflectils.FindValueOfStructFieldsByKind(v, reflect.String, reflect.Int))
	// Output: map[0:reflectils 2:<int Value>]
}
