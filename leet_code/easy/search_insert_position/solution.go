package search_insert_position

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < target && target <= nums[i] {
			return i
		}
	}
	return 0
}

// more simple
func searchInsert2(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] == target:
			return i
		case nums[i] > target:
			return i
		}
	}
	return len(nums)
}
