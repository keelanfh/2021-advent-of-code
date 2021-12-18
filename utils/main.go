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
