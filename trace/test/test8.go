package main

import (
	"fmt"
	"os"
)

func foo1() error {

	var err error
	return err
}

func foo2() error {
	var err *os.PathError
	return err
}

func main() {
	// err := foo1()
	// if err != nil {
	// 	fmt.Println("false")
	// 	return
	// }

	// fmt.Println("true")

	err := foo2()
	if err != nil {
		fmt.Println("false")
		return
	}

	fmt.Println("true")
}
