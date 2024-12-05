package calcs

func IsSafeDecreasing(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i] > nums[i+1]) && (nums[i] - nums[i+1] <= 3) {
      continue
    } else {
      return 0
    }
	}
	return 1
}

func IsSafeIncreasing(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i] < nums[i+1]) && (nums[i+1] - nums[i] <= 3) {
      continue
    } else {
      return 0
    }
	}
	return 1
}

