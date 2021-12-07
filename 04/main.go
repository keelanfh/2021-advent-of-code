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
	file, err := os.Open("04/test.txt")
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

	boardResults := make([][5][5]bool, len(bingoCards))

	for _, calledNumber := range calledNumbers {
		for i, bingoCard := range bingoCards {
			for j, row := range bingoCard {
				for k, number := range row {
					if number == calledNumber {
						boardResults[i][j][k] = true
					}
				}
			}
		}

		// Each iteration we need to check if the board has a winning row
		var won bool
		wonBoards := 0
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
				}
			}

			for i := 0; i < 5; i++ {
				total := 0
				for _, row := range boardResult {
					if row[i] {
						total++
					}
					if total == 5 {
						won = true
					}
				}

			}
			var result int
			if won {
				result = 0
				for j, row := range boardResult {
					for k, called := range row {
						if !called {
							result += bingoCards[i][j][k]
						}
					}
				}
				if wonBoards == 0 {
					fmt.Println(result * calledNumber)
				}
				wonBoards++
			}
			if wonBoards == len(bingoCards) {
				fmt.Println(result * calledNumber)
			}
		}
	}
}

// 20213 (answer given above) is *too low*.
