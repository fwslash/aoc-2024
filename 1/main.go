package main

import (
	"aoc_2024_1/utils/algorithms"
	"aoc_2024_1/utils/calcs"
	"aoc_2024_1/utils/file"
	"fmt"
)

func main() {
	arr1 := file.ReadArr("./list1.txt")
	arr2 := file.ReadArr("./list2.txt")

	n1 := len(arr1) - 1
	n2 := len(arr2) - 1

	algorithms.QuickSort(arr1, 0, n1)
	algorithms.QuickSort(arr2, 0, n2)

  diff := calcs.ArrDiff(arr1, arr2)

  fmt.Println(diff)
}
