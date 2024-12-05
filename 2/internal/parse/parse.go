package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseLine(line string) ([]int, error) {
	var nums []int

	fields := strings.Fields(line)
	for _, field := range fields {
		num, err := strconv.Atoi(field)

		if err != nil {
			fmt.Println("Error converting to integer: ", err)
			return nums, err
		}

		nums = append(nums, num)
	}

	return nums, nil
}
