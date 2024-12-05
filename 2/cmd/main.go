package main

import (
	"aoc_2024_2/internal/file"
)

func main() {
	processed, _ := file.ProcessFile("./input.txt")
	println(processed)
}
