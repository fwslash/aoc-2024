package main

import (
	"aoc_2024_3/internal/file"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	f := file.ReadFile("./cmd/input.txt")

	re := regexp.MustCompile(`(?s)don't\(\).*?(do\(\)|$)`)
	filtered := re.ReplaceAllString(f, "")

	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRe.FindAllStringSubmatch(filtered, -1)

	sum := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		product := x * y
		sum += product
	}

	fmt.Println("Total Sum:", sum)
}

