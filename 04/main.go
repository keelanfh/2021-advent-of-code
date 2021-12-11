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
	var count int
	firstWinner := true
	wonBoards := make([]bool, len(bingoCards))

	for _, calledNumber := range calledNumbers {
		count++
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

		for boardNumber, boardResult := range boardResults {
			count++
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
				for j, row := range boardResult {
					for k, called := range row {
						if !called {
							result += bingoCards[boardNumber][j][k]
						}
					}
				}
				if firstWinner {
					fmt.Println(result * calledNumber)
					firstWinner = false
				}

				// mark board as won
				wonBoards[boardNumber] = true
				fmt.Println(wonBoards)

				allBoardsWon := true
				// check if every board has been won
				for _, wonBoard := range wonBoards {
					if !wonBoard {
						allBoardsWon = false
					}
				}

				if allBoardsWon {
					fmt.Println(boardResult)
					fmt.Println(result, calledNumber)
					fmt.Println(result * calledNumber)
					return
				}
			}
		}
	}
}

// 20213 (answer given above) is *too low*.
