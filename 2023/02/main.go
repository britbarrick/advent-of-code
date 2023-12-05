package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partPtr := flag.Int("part", 1, "executes selected part")
	flag.Parse()
	// read in input file
	f, err := os.Open("data.txt")
	if err != nil {
		log.Panicf("encountered error reading file: %s", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		if *partPtr == 1 {
			part1(&total, s)
		} else if *partPtr == 2 {
			part2(&total, s)
		} else {
			log.Panicf("Invalid part %d chosen, please try again", *partPtr)
		}
	}

	fmt.Println(total)
}

func part1(total *int, s *bufio.Scanner) {
	isPossible := true
	maximums := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	// Split text input
	delim := regexp.MustCompile("[:;]+")
	parts := delim.Split(s.Text(), -1)

	for i := 1; i < len(parts); i++ {
		rounds := strings.Split(parts[i], ",")
		for j := 0; j < len(rounds); j++ {
			input := strings.Split(strings.TrimSpace(rounds[j]), " ")
			numBlocks, err := strconv.Atoi(input[0])
			if err != nil {
				log.Panicf("ERROR -- %s", err)
			}

			if  numBlocks > maximums[input[1]] {
				isPossible = false
			} 
		}
	}

	if isPossible {
		gameNumber, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			log.Panicf("ERROR - %s", err)
		}

		*total += gameNumber
	}
}

func part2(total *int, s *bufio.Scanner) {
	totals := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	delim := regexp.MustCompile("[:;]+")
	game := delim.Split(s.Text(), -1)

	for i := 1; i < len(game); i++ {
		rounds := strings.Split(game[i], ",")
		for j := 0; j < len(rounds); j++ {
			input := strings.Split(strings.TrimSpace(rounds[j]), " ")
			numBlocks, err := strconv.Atoi(input[0])
			if err != nil {
				log.Panicf("ERROR -- %s", err)
			}
			if numBlocks > totals[input[1]] {
				totals[input[1]] = numBlocks
			}
		}
	}

	*total += totals["red"] * totals["green"] * totals["blue"]
}