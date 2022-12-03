package day02

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
	draw     = 3
	win      = 6
)

var movesWLD = map[string]map[string]int{
	"A": { // rock
		"X": scissors,
		"Y": rock,
		"Z": paper,
	},
	"B": { // paper
		"X": rock,
		"Y": paper,
		"Z": scissors,
	},
	"C": { // scissors
		"X": paper,
		"Y": scissors,
		"Z": rock,
	},
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Select a scoring method: 1) XYZ for RPC 2) XYZ for WLD")
		os.Exit(1)
	}

	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("open file: %w", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	score := 0

	for s.Scan() {
		if flag.Arg(0) == "1" {
			playRoundRPC(&score, s.Text())
		} else if flag.Arg(0) == "2" {
			playRoundWLD(&score, s.Text())
		}
	}

	fmt.Println(score)
}

func playRoundRPC(score *int, move string) {
	moves := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	chars := strings.Split(move, "")
	opponentMove := chars[0]
	myMove := chars[2]
	*score += moves[myMove]

	switch opponentMove {
	case "A": // opponent plays rock
		if myMove == "X" {
			// tie
			*score += draw
		} else if myMove == "Y" {
			// win
			*score += win
		}
	case "B": // opponent plays paper
		if myMove == "Y" {
			// tie
			*score += draw
		} else if myMove == "Z" {
			// win
			*score += win
		}
	case "C": // opponent plays scissors
		if myMove == "Z" {
			// tie
			*score += draw
		} else if myMove == "X" {
			// win
			*score += win
		}
	}
}

func playRoundWLD(score *int, move string) {
	chars := strings.Split(move, "")
	wld := chars[2]
	mv := movesWLD[chars[0]][wld]
	*score += mv
	if wld == "Y" {
		*score += 3
	}
	if wld == "Z" {
		*score += 6
	}
}
