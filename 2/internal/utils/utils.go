package utils

func RemoveIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
