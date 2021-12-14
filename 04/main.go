package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkIfBoardWon(boardResult [5][5]bool) bool {
	var won bool

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
	return won
}

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
			for i, rowNumberString := range strings.Fields(scanner.Text()) {
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
	firstWinner := true
	wonBoards := make([]bool, len(bingoCards))

	// Mark the called number everywhere it appears
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

		for boardNumber, boardResult := range boardResults {
			won := checkIfBoardWon(boardResult)

			var sumUnmarkedNumbers int
			// We only want to do this the first time a board is won
			// The board might be won again later with another row
			if won && !wonBoards[boardNumber] {
				for j, row := range boardResult {
					for k, marked := range row {
						if !marked {
							sumUnmarkedNumbers += bingoCards[boardNumber][j][k]
						}
					}
				}
				// Print the result for part 1
				if firstWinner {
					fmt.Println(sumUnmarkedNumbers * calledNumber)
					firstWinner = false
				}

				// mark board as won
				wonBoards[boardNumber] = true

				// check if every board has been won
				allBoardsWon := true
				for _, wonBoard := range wonBoards {
					if !wonBoard {
						allBoardsWon = false
					}
				}

				if allBoardsWon {
					fmt.Println(sumUnmarkedNumbers * calledNumber)
					return
				}
			}
		}
	}
}
