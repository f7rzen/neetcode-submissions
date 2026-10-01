func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		pivot := left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		}
		if target > nums[pivot] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return -1
}