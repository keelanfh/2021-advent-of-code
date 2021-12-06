package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// We use uints throughout the program to make the bit manipulation cleaner

// Opens a file and returns slice of uints read from binary representation
func fileToSlice(filename string) (uint, []uint) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []uint
	var ndigits uint
	first := true
	for scanner.Scan() {
		if first {
			// Count the length of the first line and store
			ndigits = uint(len(scanner.Text()))
			first = false
		}

		// Read the binary number in
		num, _ := strconv.ParseInt(scanner.Text(), 2, 32)
		lines = append(lines, uint(num))
	}
	return ndigits, lines
}

// Returns the bit at the given index (0 is rightmost)
// This is done by shifting the number to the right by the index
// Then we mask it to get the last bit
func getBitFromInt(num uint, bit uint) uint {

	return (num >> bit) & 1
}

// Returns the most and least common digit for each position in a slice of ints
func mostLeastCommonDigit(ndigits uint, nums []uint) (uint, uint) {
	nrows := uint(len(nums))
	totalOnes := make([]uint, ndigits)

	for _, num := range nums {
		// Loop through the digit right to left
		var digit uint
		for digit = 0; digit < ndigits; digit++ {
			totalOnes[digit] += getBitFromInt(num, digit)
		}
	}

	// Now we half the length so we can see if we have more 1s than 0s
	// Increment nrows as a hack to get it to round up (7/2 = 3 but we want 4)
	nrows++
	nrows /= 2
	var gamma uint
	for i, total := range totalOnes {
		if total >= nrows {
			// We're then storing this in gamma, which we can easily print out in decimal
			// 1 << i is the same as 2^i
			gamma += 1 << i
		}
	}

	// Now we need to invert all the bits in gamma to get epsilon (least common digit)
	// This is done by doing a bitwise NOT on gamma
	// We then use a mask to get the last 12 bits
	// The mask is 2^12-1 i.e. 111111111111111
	var mask uint
	mask = (1 << ndigits) - 1
	epsilon := ^gamma & mask

	return gamma, epsilon

}

// Find rating, for part 2 of the problem
// mode is 1 for gamma (most common), 0 for epsilon (least common)
func findRating(gamma, bit, ndigits, mode uint, nums []uint) uint {
	mostCommon := getBitFromInt(gamma, bit)
	var result []uint
	for _, num := range nums {
		if getBitFromInt(num, bit) == mostCommon {
			result = append(result, num)
		}
	}
	if len(result) > 1 {
		if mode == 1 {
			gamma, _ = mostLeastCommonDigit(ndigits, result)
		} else {
			_, gamma = mostLeastCommonDigit(ndigits, result)
		}
		return findRating(gamma, bit-1, ndigits, mode, result)
	} else {
		return result[0]
	}
}

func main() {
	ndigits, nums := fileToSlice("03/input.txt")

	// Part 1
	gamma, epsilon := mostLeastCommonDigit(ndigits, nums)

	fmt.Println(gamma * epsilon)

	// Part 2
	gamma = findRating(gamma, ndigits, ndigits, 1, nums)
	epsilon = findRating(epsilon, ndigits, ndigits, 0, nums)

	fmt.Println(gamma * epsilon)
}
