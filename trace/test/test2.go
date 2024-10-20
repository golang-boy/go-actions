package main

import "fmt"

// 代码一
func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Printf("arr pointer: %p\n", &arr)
	// test(&arr)
	// fmt.Printf("arr: %v\n", arr)

	test2(&arr)
	fmt.Printf("arr: %v\n", arr)
}

func test(data *[]int) {
	fmt.Printf("data pointer: %p\n", data)
	(*data)[0] = 100
}

func test2(data *[]int) {
	fmt.Printf("data pointer: %p\n", data)
	*data = append(*data, 100)
	(*data)[0] = 100
}
