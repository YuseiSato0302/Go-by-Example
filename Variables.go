package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	f := "apple"
	fmt.Println(f)
}