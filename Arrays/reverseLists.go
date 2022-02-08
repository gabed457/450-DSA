package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4, 5, 6}
	var end int = len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if i > end {
			continue
		}
		arr[i], arr[end] = arr[end], arr[i]
		end = end - 1
	}
	fmt.Println(arr)
}
