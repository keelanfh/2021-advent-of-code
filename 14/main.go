package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var polymerTemplate []rune
	first := true
	re := regexp.MustCompile(`^([A-Z]{2}) -> ([A-Z])$`)
	findToInsert := make(map[string]rune)

	for scanner.Scan() {
		line := scanner.Text()
		if first {
			polymerTemplate = []rune(line)
			first = false
		} else if line != "" {
			matches := re.FindStringSubmatch(line)
			findToInsert[matches[1]] = []rune(matches[2])[0]
		}
	}
	var output []rune
	for step := 1; step <= 10; step++ {
		output = make([]rune, 0)

		for i := 0; i <= len(polymerTemplate)-2; i++ {

			output = append(output, polymerTemplate[i])
			insertString := findToInsert[string(polymerTemplate[i:i+2])]

			// 0 is the nil value for rune, because rune is just int32
			if insertString != 0 {
				output = append(output, insertString)
			}

		}

		output = append(output, polymerTemplate[len(polymerTemplate)-1])

		polymerTemplate = output
		fmt.Println(step, len(polymerTemplate))

	}

	counter := make(map[rune]int)

	for _, char := range output {
		counter[char]++
	}

	minValue := math.MaxInt32
	maxValue := 0

	for _, value := range counter {
		if value > maxValue {
			maxValue = value
		}
		if value < minValue {
			minValue = value
		}
	}

	fmt.Println(maxValue - minValue)

}
