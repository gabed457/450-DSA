package max_min_from_list

func maxMinArr(list []int) []int {
	var min = list[0]
	var max = list[0]
	for i := 0; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
		if list[i] > max {
			max = list[i]
		}
	}
	return []int{min, max}
}
