package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		panic("Provide argument specifying which part to execute")
	}

	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("open file: %w", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	priority := 0

	if flag.Arg(0) == "1" {
		for s.Scan() {
			findPackPriority(&priority, s.Text())
		}
	} else if flag.Arg(0) == "2" {
		findGroupPriority(&priority, s)
	}

	fmt.Println(priority)
}

func findGroupPriority(p *int, s *bufio.Scanner) {
	more := true
	for more {
		group := []string{}
		for i := 0; i < 3; i++ {
			more = s.Scan()
			group = append(group, s.Text())
		}
		for _, r := range group[0] {
			if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
				calculatePriority(p, r)
				break
			}
		}
	}
}

func findPackPriority(p *int, sack string) {
	p1 := sack[0:(len(sack) / 2)]
	p2 := sack[(len(sack) / 2):]

	for _, r := range p1 {
		if strings.ContainsRune(p2, r) {
			calculatePriority(p, r)
		}
	}
}

func calculatePriority(p *int, r rune) {
	priority := int(r) - 96
	if priority < 0 {
		priority = int(r) - 38
	}
	*p += priority
}
