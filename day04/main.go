package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var i = 0

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
	overlap := 0

	for s.Scan() {
		if flag.Arg(0) == "1" {
			findContainment(&overlap, s.Text())
		} else if flag.Arg(0) == "2" {
			findOverlap(&overlap, s.Text())
		}
	}

	fmt.Println(overlap)
}

func findContainment(counter *int, pair string) {
	a1, a2 := splitAssignments(pair)

	if compareForContainment(a1, a2) || compareForContainment(a2, a1) {
		*counter++
	}
}

func findOverlap(counter *int, pair string) {
	a1, a2 := splitAssignments(pair)

	if compareForOverlap(a1, a2) || compareForOverlap(a2, a1) {
		*counter++
	}
}

func splitAssignments(pair string) ([]string, []string) {
	assignments := strings.Split(pair, ",")
	assignment1, assignment2 := assignments[0], assignments[1]
	a1 := strings.Split(assignment1, "-")
	a2 := strings.Split(assignment2, "-")

	return a1, a2
}

func compareForContainment(a, b []string) bool {
	aStart, _ := strconv.Atoi(a[0])
	aEnd, _ := strconv.Atoi(a[1])
	bStart, _ := strconv.Atoi(b[0])
	bEnd, _ := strconv.Atoi(b[1])
	if aStart >= bStart && aStart <= bEnd && aEnd >= bStart && aEnd <= bEnd {
		return true
	}

	return false
}

func compareForOverlap(a, b []string) bool {
	aStart, _ := strconv.Atoi(a[0])
	aEnd, _ := strconv.Atoi(a[1])
	bStart, _ := strconv.Atoi(b[0])
	bEnd, _ := strconv.Atoi(b[1])
	if (aStart >= bStart && aStart <= bEnd) || (aEnd >= bStart && aEnd <= bEnd) {
		return true
	}

	return false
}
