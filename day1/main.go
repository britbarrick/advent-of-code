package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// read in file
	f, err := os.Open("data.txt")
	if err != nil {
		log.Panicf("encountered error reading file: %s", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	elf, one, two, three := 0, 0, 0, 0

	for s.Scan() {
		// split data at empty lines
		if s.Text() == "" {
			if elf > one {
				three = two
				two = one
				one = elf
			} else if elf > two {
				three = two
				two = elf
			} else if elf > three {
				three = elf
			}

			elf = 0
		} else {
			// sum all data per elf
			cal, err := strconv.Atoi(s.Text())
			if err != nil {
				log.Default().Printf("Cannot convert %s to int", s.Text())
			}

			elf += cal
		}
	}

	total := one + two + three

	fmt.Printf("The top three have: %d, %d, %d\nFor a total of: %d\n",
		one, two, three, total)
}
