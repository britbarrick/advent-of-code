package day05

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stacksList []string

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

	stacks := stacksList{}

	s := bufio.NewScanner(f)
	for i := 0; i < 10; i++ {
		s.Scan()
		stacks = append(stacks, s.Text())
	}

	stackMap := stacks.convert()
	for s.Scan() {
		if flag.Arg(0) == "1" {
			stackMap = processStepOne(stackMap, s.Text())
		} else if flag.Arg(0) == "2" {
			stackMap = processStepTwo(stackMap, s.Text())
		} else {
			panic("Invalid option selected, choose 1) Part 1 -OR- 2) Part 2")
		}
	}

	fmt.Println(findOutput(stackMap))
}

func (s stacksList) convert() map[int][]string {
	stackMap := map[int][]string{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	for k, stack := range s {
		if k < 8 {
			j := 1
			for i := 1; i < len(stack); i += 4 {
				if string(stack[i]) != " " {
					stackMap[j] = append(stackMap[j], string(stack[i]))
				}
				j++
			}
		}
	}

	return stackMap
}

func processStepOne(stack map[int][]string, step string) map[int][]string {
	s := strings.Split(step, " ")
	numSteps, _ := strconv.Atoi(s[1])
	origin, _ := strconv.Atoi(s[3])
	destination, _ := strconv.Atoi(s[5])

	for numSteps > 0 {
		moving := stack[origin][0]
		stack[origin] = stack[origin][1:]
		stack[destination] = append([]string{moving}, stack[destination]...)
		numSteps--
	}

	return stack
}

func processStepTwo(stack map[int][]string, step string) map[int][]string {
	s := strings.Split(step, " ")
	numCrates, _ := strconv.Atoi(s[1])
	origin, _ := strconv.Atoi(s[3])
	destination, _ := strconv.Atoi(s[5])

	moving := stack[origin][0:numCrates]
	newOrigin := []string{}
	newOrigin = append(newOrigin, stack[origin][numCrates:]...)
	newDestination := append(moving, stack[destination]...)

	stack[origin] = newOrigin
	stack[destination] = newDestination

	return stack
}

func findOutput(stacks map[int][]string) string {
	out := ""
	for i := 1; i <= len(stacks); i++ {
		out += (stacks[i])[0]
	}

	return out
}
