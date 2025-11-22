package task1

func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int)
	for k, v := range nums {
		j, ok := map1[target-v]
		if ok {
			return []int{j, k}
		}
		map1[v] = k
	}
	return nil
}
