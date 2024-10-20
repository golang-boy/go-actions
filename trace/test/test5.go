package main

import "fmt"

func main() {

	defer catch("main") // 恢复后，从b函数开始捕获panic

	a()

}

func a() {
	defer b()
	panic("a panic")
}

func b() {
	defer fb()
	panic("b panic")
}

func fb() {
	defer catch("fb") // 只会捕捉，当前函数，以及当前函数调用函数的panic
	panic("fb panic")
}

func catch(name string) {
	if r := recover(); r != nil {
		fmt.Println(name, "recovered:", r)
	}
}

//  (base) PS C:\Users\86186\workspace\go-actions\go-actions\trace\test> go run .\test5.go
//  fb recovered: fb panic
//  main recovered: b panic
