package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

type SortableRuneSlice []rune

func (a SortableRuneSlice) Len() int           { return len(a) }
func (a SortableRuneSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableRuneSlice) Less(i, j int) bool { return a[i] < a[j] }

func intersection(first, second []rune) []rune {
	// Order doesn't matter for this one
	var result []rune
	for _, char := range first {
		for _, char2 := range second {
			if char == char2 {
				result = append(result, char)
				break
			}
		}
	}
	return result
}

func difference(first, second []rune) []rune {
	// Order matters - this is like subtraction!
	var result []rune
	for _, char := range first {
		var match bool
		for _, char2 := range second {
			if char == char2 {
				match = true
				break
			}
		}
		if !match {
			result = append(result, char)
		}
	}
	return result
}

func main() {
	scanner := utils.ReadFileLines("08/input.txt")

	var total int

	segmentsToDigits := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	var result int

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " | ")
		outputDigitsSlice := strings.Split(splitLine[1], " ")

		for _, outputDigit := range outputDigitsSlice {
			switch len(outputDigit) {
			case 2, 3, 4, 7:
				total++
			}
		}

		signalPatternsSlice := strings.Split(splitLine[0], " ")
		var lengthToSignalPatterns [8][][]rune
		segmentsToSignalPatterns := make(map[rune]rune)

		for _, signalPattern := range signalPatternsSlice {
			lengthToSignalPatterns[len(signalPattern)] = append(lengthToSignalPatterns[len(signalPattern)], []rune(signalPattern))
		}

		// deduce mapping of segment a
		// difference of digits 1 and 7, i.e. length 2 and 3
		segmentsToSignalPatterns['a'] = []rune(difference(lengthToSignalPatterns[3][0], lengthToSignalPatterns[2][0]))[0]

		diff41 := difference(lengthToSignalPatterns[4][0], lengthToSignalPatterns[2][0])

		intersection532 := intersection(intersection(lengthToSignalPatterns[5][0], lengthToSignalPatterns[5][1]), lengthToSignalPatterns[5][2])

		segmentsToSignalPatterns['d'] = intersection(diff41, intersection532)[0]

		segmentsToSignalPatterns['b'] = difference(diff41, []rune{segmentsToSignalPatterns['d']})[0]

		segmentsToSignalPatterns['g'] = difference(intersection532, []rune{segmentsToSignalPatterns['d'], segmentsToSignalPatterns['a']})[0]

		intersection069 := intersection(intersection(lengthToSignalPatterns[6][0], lengthToSignalPatterns[6][1]), lengthToSignalPatterns[6][2])

		segmentsToSignalPatterns['f'] = difference(intersection069, []rune{segmentsToSignalPatterns['a'], segmentsToSignalPatterns['b'], segmentsToSignalPatterns['g']})[0]

		segmentsToSignalPatterns['c'] = difference(lengthToSignalPatterns[2][0], []rune{segmentsToSignalPatterns['f']})[0]

		signalPatternsToSegments := make(map[rune]rune)

		for k, v := range segmentsToSignalPatterns {
			signalPatternsToSegments[v] = k
		}

		for i := 'a'; i <= 'g'; i++ {
			if signalPatternsToSegments[i] == 0 {
				signalPatternsToSegments[i] = 'e'
			}
		}

		// target mapping for small test is
		// a:d
		// b:e
		// c:a
		// d:f
		// e:g
		// f:b
		// g:c

		// now look at the output strings
		multiple := 1000
		for _, outputString := range outputDigitsSlice {
			var segments SortableRuneSlice
			for _, char := range outputString {
				segment := signalPatternsToSegments[char]
				segments = append(segments, segment)
			}
			sort.Sort(segments)
			result += segmentsToDigits[string(segments)] * multiple
			multiple /= 10
		}
	}

	fmt.Println(total)
	fmt.Println(result)

}
