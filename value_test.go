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

func ExampleNewSliceOfValue() {
	v := reflectils.NewSliceOfValue("reflectils", reflect.ValueOf(1), 2.3)
	_ = reflectils.IterValue(v, func(i int, e reflect.Value) (err error) {
		fmt.Println(e)
		return nil
	})
	// Output:
	// reflectils
	// 1
	// 2.3
}

func ExampleIterValue() {
	v := reflectils.NewSliceOfValue("reflectils", reflect.ValueOf(1), 2.3)
	_ = reflectils.IterValue(v, func(i int, e reflect.Value) (err error) {
		if e.Kind() == reflect.Int {
			err = reflectils.ErrStopIteration
		}
		fmt.Println(i, e)
		return
	})
	// Output:
	// 0 reflectils
	// 1 1
}
