package ptr

import (
	"fmt"
)

func ExampleTo() {
	type dummy struct {
		IntPtr *int
	}

	d := dummy{
		IntPtr: To(123),
	}
	fmt.Printf("%T", d.IntPtr)

	// Output:
	// *int
}

func ExampleValue() {
	i := 123
	ptr := &i

	value := Value(ptr)
	fmt.Println(value)

	s := "abc"
	str := Value(&s)
	fmt.Println(str)

	// Output:
	// 123
	// abc
}
