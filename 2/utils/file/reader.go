package file

import (
	"aoc_2024_2/utils/calcs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Reader(filePath string) int {
	safeCount := 0

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error occured during reading file %s", err)
		return safeCount
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var nums []int
		line := scanner.Text()
		fields := strings.Fields(line)

		for _, field := range fields {
			num, err := strconv.Atoi(field)

			if err != nil {
				fmt.Println("Error converting to integer: ", err)
				return safeCount
			}

			nums = append(nums, num)
		}

		if nums[0] < nums[1] {
			safeCount += calcs.IsSafeIncreasing(nums)
		} else {
			safeCount += calcs.IsSafeDecreasing(nums)
		}

	}

	return safeCount
}
