package reflectils_test

import (
	"fmt"

	"github.com/acarsync/reflectils"
)

func ExampleStringFromBytes() {
	fmt.Println(reflectils.StringFromBytes([]byte("Hello World")))
	// Output:
	// Hello World
}

func ExampleStringToBytes() {
	fmt.Println(reflectils.StringFromBytes(reflectils.StringToBytes("Hello World")))
	// Output:
	// Hello World
}
