package task1

func removeDuplicates(nums []int) int {
	var k1 int = 0
	var k2 int = 1
	for k2 < len(nums) {
		if nums[k1] == nums[k2] {
			k2++
		} else {
			k1++
			nums[k1] = nums[k2]
		}
	}
	return k1 + 1
}
