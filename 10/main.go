package main

import (
	"container/ring"
	"fmt"
	"sort"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func isClosing(c rune) bool {
	switch c {
	case ')', ']', '}', '>':
		return true
	default:
		return false
	}
}

func printRing(r *ring.Ring) {
	for j := 0; j < r.Len(); j++ {
		if r.Value == nil {
			fmt.Print("NIL")
		} else {
			// .(rune) is a type assertion
			fmt.Print(string(r.Value.(rune)))
		}
		r = r.Next()
	}
	fmt.Println()
}

func main() {
	scanner := utils.ReadFileLines("10/input.txt")

	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	values := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	// we value these based on the opening, rather than closing
	// this just keeps it simpler
	completionValues := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	var syntaxErrorScore int
	var corrupted bool
	var autocompleteScores []int

	for scanner.Scan() {
		line := scanner.Text()
		corrupted = false

		// Set up the ring
		// We add an extra element
		// So we can keep track of the beginning/end
		r := ring.New(len(line) + 1)

		// Load the characters in
		for _, c := range line {
			r.Value = c
			r = r.Next()
		}

		// Move back to the beginning of the ring
		r = r.Next()

		for r.Value != nil {
			if isClosing(r.Value.(rune)) {
				if r.Prev().Value.(rune) == match[r.Value.(rune)] {
					// remove these from the ring
					// r.Unlink removes the next n, so we need to move back 2
					// to remove the last 2
					r = r.Prev().Prev()
					r.Unlink(2)
				} else {
					syntaxErrorScore += values[r.Value.(rune)]
					corrupted = true
					// Only interested in the first incorrect closing character
					// on each line, so break out of this for loop
					break
				}
			}
			r = r.Next()

		}
		// We only need to check autocomplete if the line isn't corrupted
		if !corrupted {

			// we're going to work through the ring backwards
			// because we need to close brackets in reverse order

			// first move back one position from the nil
			r = r.Prev()
			score := 0
			for r.Value != nil {
				score *= 5
				score += completionValues[r.Value.(rune)]
				r = r.Prev()
			}

			autocompleteScores = append(autocompleteScores, score)
		}
	}

	// Part 1 output
	fmt.Println(syntaxErrorScore)

	// Part 2 output
	sort.Ints(autocompleteScores)
	fmt.Println(autocompleteScores[len(autocompleteScores)/2])
}
