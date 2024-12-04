package algorithms

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	i := start - 1

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, i+1, end)
	return i + 1
}

func swap(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}

func QuickSort(arr []int, start, end int) {
	if start < end {
		pi := partition(arr, start, end)

		QuickSort(arr, pi+1, end)
		QuickSort(arr, start, pi-1)
	}
}
