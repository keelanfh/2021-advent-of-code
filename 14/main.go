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

	var polymerTemplate string
	first := true
	re := regexp.MustCompile(`^([A-Z])([A-Z]) -> ([A-Z])$`)
	findToReplace := make(map[string]rune)

	for scanner.Scan() {
		line := scanner.Text()
		if first {
			polymerTemplate = line
			first = false
		} else if line != "" {
			matches := re.FindStringSubmatch(line)

			findString := matches[1] + matches[2]
			replaceString := []rune(matches[3])[0]

			findToReplace[findString] = replaceString
		}
	}
	polymerTemplateRunes := []rune(polymerTemplate)
	var resultString []rune
	for step := 1; step <= 40; step++ {
		resultString = make([]rune, 0)

		for i := 0; i <= len(polymerTemplateRunes)-2; i++ {

			resultString = append(resultString, polymerTemplateRunes[i])
			insertString := findToReplace[string(polymerTemplateRunes[i:i+2])]
			if insertString != 0 {
				resultString = append(resultString, insertString)
			}

			// fmt.Println("result:", string(resultString))
		}

		resultString = append(resultString, polymerTemplateRunes[len(polymerTemplateRunes)-1])

		// fmt.Println("step", step+1, string(resultString))

		polymerTemplateRunes = resultString
		fmt.Println(step, len(polymerTemplateRunes))

	}

	counter := make(map[rune]int)

	for _, char := range resultString {
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
