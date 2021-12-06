package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// File opening
	file, err := os.Open("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1, and putting everything in a slice for Part 2
	var nums []int

	scanner := bufio.NewScanner(file)
	current := math.MaxInt32
	var previous int
	increasing := 0
	for scanner.Scan() {
		previous = current
		current, err = strconv.Atoi(scanner.Text())
		nums = append(nums, current)
		if current > previous {
			increasing++
		}
	}

	fmt.Println(increasing)

	// Part 2

	// Take a three period moving average
	var sum int
	var sums []int
	for i := 0; i < len(nums)-2; i++ {
		sum = 0
		for j := 0; j < 3; j++ {
			sum += nums[i+j]
		}
		sums = append(sums, sum)
	}

	previous = math.MaxInt32
	increasing = 0
	for _, sum := range sums {
		if sum > previous {
			increasing++
		}
		previous = sum
	}

	fmt.Println(increasing)
}
