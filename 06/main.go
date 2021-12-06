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
	ageMap := make(map[int]int)
	for _, s := range sliceOfStrings {
		i, _ := strconv.Atoi(s)
		ageMap[i]++
	}

	var newAgeMap map[int]int

	for day := 1; day <= 256; day++ {
		newAgeMap = make(map[int]int)
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

	result := 0
	for _, v := range ageMap {
		result += v
	}
	fmt.Println(result)

}
