package file

import (
	"aoc_2024_2/internal/parse"
	"aoc_2024_2/internal/sequance"
	"bufio"
	"fmt"
	"os"
)

func ProcessFile(filePath string) (int, error) {
	safeCount := 0
	file, err := os.Open(filePath)

	if err != nil {
		return safeCount, fmt.Errorf("Error occured during reading file %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var nums []int
		line := scanner.Text()
		nums, err = parse.ParseLine(line)

		if err != nil {
			return safeCount, fmt.Errorf("Error occured during parseLine: %s", err)
		}

		safeCount += sequance.ProcessSequance(nums)
	}

	return safeCount, nil
}
