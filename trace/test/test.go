package main

import "fmt"

func main() {

	var a uint8 = 255

	var b uint8 = 1

	fmt.Println(a + b)

	F()
	temp := make([]int, 0, 10)
	F1(temp)

	F2 := func() {
		// fmt.Println("hello")
		var a int
		temp := make([]int, 0, 10)
		a = 1
		temp[0] = a
	}
	// F2()

	// fmt.Printf("%p\n", F2)

	F3 := F2
	F3()

	F2()
}

func F() {
	temp := make([]int, 0, 10)

	fmt.Println(temp)
}

func F1(temp []int) {

	a := temp
	fmt.Println(a)
}

// 逃逸:
//  1. 函数内分配指针，作为返回值返回
//  2. 栈空间不足
//  3. interface,入参为
//  4. slice与map中是指针元素
//  5. 与闭包内的变量相关

//   (base) PS C:\Users\86186\workspace\go-actions\go-actions\trace> go build -gcflags "-m -l" .\test.go
//   # command-line-arguments
//   ./test.go:20:14: make([]int, 0, 10) escapes to heap
//   ./test.go:22:13: ... argument does not escape
//   ./test.go:22:14: temp escapes to heap
//   ./test.go:25:9: leaking param: temp
//   ./test.go:28:13: ... argument does not escape
//   ./test.go:28:14: a escapes to heap
//   ./test.go:11:13: ... argument does not escape
//   ./test.go:11:16: a + b escapes to heap
//   ./test.go:14:14: make([]int, 0, 10) escapes to heap
