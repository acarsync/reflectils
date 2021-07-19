package reflectils_test

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/acarsync/reflectils"
)

type ExampleStruct struct {
	A string
	B uint
	C int
}

func ExampleMapStructFields() {
	v := reflect.TypeOf((*ExampleStruct)(nil)).Elem()
	reflectils.MapStructFields(v, func(i int, v reflect.StructField) error {
		fmt.Println(i, v.Name, v.Type)
		return nil
	})
	// Output:
	// 0 A string
	// 1 B uint
	// 2 C int
}

func ExampleMapTypeOfStructFields() {
	v := reflect.TypeOf((*ExampleStruct)(nil)).Elem()
	reflectils.MapTypeOfStructFields(v, func(i int, v reflect.Type) error {
		fmt.Println(i, v)
		return nil
	})
	// Output:
	// 0 string
	// 1 uint
	// 2 int
}

func ExampleMapTypeOfStructFields_withErrorHandling() {
	v := reflect.TypeOf((*ExampleStruct)(nil)).Elem()
	err := reflectils.MapTypeOfStructFields(v, func(i int, v reflect.Type) error {
		switch v.Kind() {
		default:
			return errors.New("unexpected")
		case reflect.String, reflect.Uint:
		}
		fmt.Println(i, v)
		return nil
	})
	fmt.Println(err)
	// Output:
	// 0 string
	// 1 uint
	// unexpected
}

func ExampleMapValueOfStructFields() {
	v := reflect.ValueOf(ExampleStruct{
		A: "reflectils",
		B: 123,
		C: 456,
	})
	reflectils.MapValueOfStructFields(v, func(i int, v reflect.Value) error {
		fmt.Println(i, v)
		return nil
	})
	// Output:
	// 0 reflectils
	// 1 123
	// 2 456
}

func ExampleMapValueOfStructFields_withErrorHandling() {
	v := reflect.ValueOf(ExampleStruct{
		A: "reflectils",
		B: 123,
		C: 456,
	})
	err := reflectils.MapValueOfStructFields(v, func(i int, v reflect.Value) error {
		switch v.Kind() {
		default:
			return errors.New("unexpected")
		case reflect.String, reflect.Uint:
		}
		fmt.Println(i, v)
		return nil
	})
	fmt.Println(err)
	// Output:
	// 0 reflectils
	// 1 123
	// unexpected
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
