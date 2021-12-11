package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

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
