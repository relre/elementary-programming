package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(arr []int, size int) {
	var slice []int
	if size <= 0 {
		fmt.Println()
		return
	}
	result := make([][]int, 0, len(arr)/size+1)
	for len(arr) >= size {
		slice, arr = arr[:size], arr[size:]
		result = append(result, slice)
	}
	if len(arr) > 0 {
		result = append(result, arr[:])
	}
	fmt.Println(result)
}
