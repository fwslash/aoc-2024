package calcs

func ArrDiff(arr1 []int, arr2 []int) int {
	diff := 0

	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			diff += arr1[i] - arr2[i]
		} else {
			diff += arr2[i] - arr1[i]
		}
	}

	return diff
}
