package task1

func singleNumber(nums []int) int {
	map1 := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		map1[nums[i]] = map1[nums[i]] + 1
	}
	for k, v := range map1 {
		if v == 1 {
			return k
		}
	}
	return -1
}
