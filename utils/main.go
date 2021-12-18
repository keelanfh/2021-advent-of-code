package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(path string) *bufio.Scanner {
	file, err := os.Open("02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}
