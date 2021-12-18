package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
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
