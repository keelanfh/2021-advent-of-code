package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {

	scanner := utils.ReadFileLines("06/input.txt")

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	sliceOfStrings := strings.Split(line, ",")
	var ageMap [9]int
	for _, s := range sliceOfStrings {
		i, _ := strconv.Atoi(s)
		ageMap[i]++
	}

	for day := 1; day <= 256; day++ {
		var newAgeMap [9]int
		for k, v := range ageMap {
			if k == 0 {
				newAgeMap[6] += v
				newAgeMap[8] += v
			} else {
				newAgeMap[k-1] += v
			}
		}
		ageMap = newAgeMap

	}

	var result int
	for _, v := range ageMap {
		result += v
	}
	fmt.Println(result)

}
