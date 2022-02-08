package reverse_lists

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
