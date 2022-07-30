func searchInsert(nums []int, target int) int {
    l, h := 0, len(nums)-1
    if target > nums[h] {
        return len(nums)
    }
    for l != h {
        var m int = (l + h)/2
        if nums[m] >= target {
            h = m
        } else {
            l = m + 1
        }
    }
    return l
}
