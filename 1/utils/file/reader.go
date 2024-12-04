package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadArr(filePath string) []int {
	var numbers []int
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error occured during reading file")
		return numbers
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Error occured during scan")
			return numbers
		}

		numbers = append(numbers, num)
	}

	return numbers
}
