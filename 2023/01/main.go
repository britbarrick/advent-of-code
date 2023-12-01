package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	num1, num2 := "", "";
	for _, ch := range s.Text() {
		if unicode.IsNumber(ch) && num1 == "" {
			num1 = string(ch)
			num2 = string(ch)
		} else if unicode.IsNumber(ch) {
			num2 = string(ch)
		}
	}

	calibration, err := strconv.ParseInt(fmt.Sprintf("%s%s", num1, num2), 10, 64)
	if err != nil {
		log.Panic("Conversion issue!!")
	}
	
	*total += int(calibration)
}

func part2(total *int, s *bufio.Scanner) {
	line := s.Text()
	numbers:= map[string]string{
		"one": "o1e",
		"two": "t2o",
		"three": "th3ee",
		"four": "fo4ur",
		"five": "fi5ve",
		"six": "s6x",
		"seven": "se7en",
		"eight": "ei8ht",
		"nine": "ni9ne",
	}

	num1, num2 := "", "";
	for key, val := range numbers {
		if strings.Contains(line, key) {
			line = strings.ReplaceAll(line, key, fmt.Sprint(val))
		}
	}
	for _, ch := range line {
		if unicode.IsNumber(ch) && num1 == "" {
			num1 = string(ch)
			num2 = string(ch)
		} else if unicode.IsNumber(ch) {
			num2 = string(ch)
		}
	}

	calibration, err := strconv.ParseInt(fmt.Sprintf("%s%s", num1, num2), 10, 64)
	if err != nil {
		log.Panic("Conversion issue!!")
	}
	
	*total += int(calibration)
}