package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {

	var hpos, depth int

	scanner := utils.ReadFileLines("02/input.txt")
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		op, argString := split[0], split[1]
		arg, _ := strconv.Atoi(argString)
		switch op {
		case "forward":
			hpos += arg
		case "down":
			depth += arg
		case "up":
			depth -= arg
		}
	}

	fmt.Println(hpos * depth)

	scanner = utils.ReadFileLines("02/input.txt")

	hpos, depth = 0, 0
	var aim int
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		op, argString := split[0], split[1]
		arg, _ := strconv.Atoi(argString)
		switch op {
		case "forward":
			hpos += arg
			depth += aim * arg
		case "down":
			aim += arg
		case "up":
			aim -= arg
		}
	}

	fmt.Println(hpos * depth)
}
