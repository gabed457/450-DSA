package main

import "fmt"

func reverseList(list []int) []int {
	var end int = len(list) - 1
	for i := 0; i < len(list); i++ {
		if i > end {
			break
		}
		list[i], list[end] = list[end], list[i]
		end = end - 1
	}
	return list
}
func main() {
	//Test reverseList Func with array
	arr := [...]int{1, 4, 7, 22, 135}
	fmt.Println(reverseList(arr[:]))
	//Test reverseList Func with array
	mySlice := []int{-3, 3, 4, -5, 6000}
	mySlice = append(mySlice, 12)
	fmt.Println(reverseList(mySlice))
}
