package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " | ")
		// signalPatternsString := splitLine[0]
		// signalPatternsSlice := strings.Split(signalPatternsString, " ")

		outputDigitsString := splitLine[1]
		outputDigitsSlice := strings.Split(outputDigitsString, " ")

		for _, outputDigit := range outputDigitsSlice {
			switch len(outputDigit) {
			case 2, 3, 4, 7:
				total++
			}
		}
	}
	fmt.Println(total)
}
