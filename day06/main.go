package day06

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	s.Scan()
	input := s.Text()

	if flag.Arg(0) == "1" {
		marker := make([]byte, 4, 4)
		for i := 0; i < len(input); i++ {
			newChar := input[i]
			if byteSliceContains(newChar, marker) || matches(marker) {
				for j := 1; j < 4; j++ {
					marker[j-1] = marker[j]
				}
				marker[3] = newChar
			} else if !matches(marker) {
				fmt.Println("MARKER", i-1)
				break
			}
		}
	} else if flag.Arg(0) == "2" {
		for i := 14; i < len(input); i++ {
			message := map[rune]struct{}{}
			for _, r := range input[i-14 : i] {
				message[r] = struct{}{}
			}
			if len(message) == 14 {
				fmt.Println("MESSAGE", i)
			}
		}
	} else {
		panic("Provide valid choice 1) Marker -OR- 2) Message")
	}
}

func byteSliceContains(b byte, s []byte) bool {
	for _, v := range s {
		if b == v {
			return true
		}
	}

	return false
}

func matches(s []byte) bool {
	dict := make(map[byte]int)
	for _, v := range s {
		dict[v] = dict[v] + 1
		if dict[v] == 2 {
			return true
		}
	}
	return false
}
