package main

import (
	"fmt"
)

func main() {

	// a := 1
	// defer func(b int) {
	// 	fmt.Println(b)
	// }(a)

	// a = 99

	i := f()
	fmt.Printf("main: i=%d, g = %d\n", i, g)

}

var g = 100

func f() (r int) {

	r = g
	defer func() {
		r = 200
	}()

	r = 0
	return r
}
