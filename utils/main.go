package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileLines(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func ReadFileGridInts(path string) [][]int {
	var grid [][]int
	scanner := ReadFileLines(path)
	for scanner.Scan() {
		var list []int
		line := scanner.Text()
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			list = append(list, num)
		}
		grid = append(grid, list)
	}
	return grid
}

type StringSlice []string

func (sa StringSlice) Contains(s string) bool {
	for _, t := range sa {
		if s == t {
			return true
		}
	}
	return false
}

func StringIsLower(s string) bool {
	return strings.ToLower(s) == s
}

// func PrintGrid(grid [][]interface{}) {
// 	for _, line := range grid {
// 		for _, number := range line {
// 			if number == 0 {
// 				fmt.Print(".")
// 			} else {
// 				fmt.Print(number)
// 			}

// 		}
// 		fmt.Println()
// 	}
// }
