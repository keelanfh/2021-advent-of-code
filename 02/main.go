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
	// Part 1
	file, err := os.Open("02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	hpos := 0
	depth := 0
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

	// Part 2
	file, err = os.Open("02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)

	hpos = 0
	depth = 0
	aim := 0
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
