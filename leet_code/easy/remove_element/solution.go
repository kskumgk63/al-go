package remove_element

// retruting value is correct, but it violate the rule.
func removeElement(nums []int, val int) int {
	n := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			n = append(n, nums[i])
		}
	}
	return len(n)
}

// Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.
func removeElement2(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i]-val == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			i-- // to handle change of array length
		}
	}
	return len(nums)
}
