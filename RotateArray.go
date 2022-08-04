func rotate(nums []int, k int) {
    k =  k % len(nums)

    rev(nums, 0, len(nums)-1)
    rev(nums, 0, k-1)
    rev(nums, k, len(nums)-1)

}

func rev(nums []int, start int, end int) {
    for start <= end {
        temp := nums[start]
        nums[start] = nums[end]
        nums[end] = temp
        end--
        start++
    }
}
