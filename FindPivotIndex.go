func pivotIndex(nums []int) int {
	for i, _ := range nums {
		if i == len(nums) {
			return -1
		}
		if sumSlice(nums[:i]) == sumSlice(nums[i+1:]) {
			return i
		}
	}
	return -1
}

func sumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}