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
	file, err := os.Open("04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var calledNumbers []int
	var bingoCards [][5][5]int
	var row int
	currentCard := -1

	first := true
	for scanner.Scan() {
		if first {
			first = false

			calledNumbersList := strings.Split(scanner.Text(), ",")
			for _, calledNumberString := range calledNumbersList {
				calledNumber, err := strconv.Atoi(calledNumberString)
				if err != nil {
					log.Fatal(err)
				}
				calledNumbers = append(calledNumbers, calledNumber)
			}
		} else if scanner.Text() == "" {
			row = 0
			currentCard++
			bingoCards = append(bingoCards, [5][5]int{})
		} else {
			rowNumberStrings := strings.Fields(scanner.Text())
			for i, rowNumberString := range rowNumberStrings {
				// fmt.Println(rowNumberString)
				number, err := strconv.Atoi(rowNumberString)
				if err != nil {
					log.Fatal(err)
				}
				bingoCards[currentCard][row][i] = number
			}
			row++
		}
	}

	var boardResults [][5][5]bool

	for _, calledNumber := range calledNumbers {
		for i, bingoCard := range bingoCards {
			boardResults = append(boardResults, [5][5]bool{})
			for j, row := range bingoCard {
				for k, number := range row {
					if number == calledNumber {
						boardResults[i][j][k] = true
					}
				}
			}
		}

		var won bool
		for i, boardResult := range boardResults {
			for _, row := range boardResult {
				total := 0
				for _, called := range row {
					if called {
						total++
					}
				}
				if total == 5 {
					won = true
					break
				}
			}
			if won {
				result := 0
				for j, row := range boardResult {
					for k, called := range row {
						if !called {
							result += bingoCards[i][j][k]
						}
					}
				}
				fmt.Println(result * calledNumber)
				return
			}
		}
	}
}
