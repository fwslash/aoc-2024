package sequance

import "aoc_2024_2/internal/utils"

func isSafeDecreasing(nums []int, popped bool) int {
	for i := len(nums) - 2; i >= 0; i-- {
		if !((nums[i] > nums[i+1]) && (nums[i]-nums[i+1] <= 3)) {
			if popped {
				return 0
			}

			if nums[i] <= nums[i+1] {
				popped = true
				nums = utils.RemoveIndex(nums, i+1)
				i--
				continue
			}

			return 0

		}

	}
	return 1
}

func isSafeIncreasing(nums []int, popped bool) int {
	for i := len(nums) - 2; i >= 0; i-- {

		if !((nums[i] < nums[i+1]) && (nums[i+1]-nums[i] <= 3)) {
			if popped {
				return 0
			}
			if nums[i] >= nums[i+1] {
				popped = true
				nums = utils.RemoveIndex(nums, i+1)
				i--
				continue
			}
			return 0
		}
	}
	return 1
}

func ProcessSequance(nums []int) int {
	popped := false

	if nums[0] == nums[1] {
		nums = utils.RemoveIndex(nums, 0)
		popped = true
	}

	if nums[0] > nums[1] {
		return isSafeDecreasing(nums, popped)
	}

	return isSafeIncreasing(nums, popped)
}
