package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	Test(PartOne(ParseInput("input_example.txt")), 161)
	fmt.Println("Part One:", PartOne(ParseInput("input.txt")))

	// Part 2
	Test(PartTwo(ParseInput("input_example2.txt")), 48)
	fmt.Println("Part Two:", PartTwo(ParseInput("input.txt")))
}

func PartOne(input string) int {
	ans := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, ins := range matches {
		m, _ := strconv.Atoi(ins[1])
		n, _ := strconv.Atoi(ins[2])
		ans += m * n
	}

	return ans
}

func PartTwo(input string) int {
	ans := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	doMult := true
	for _, ins := range matches {
		switch ins[0] {
		case "do()":
			doMult = true
		case "don't()":
			doMult = false
		default:
			if doMult {
				m, _ := strconv.Atoi(ins[1])
				n, _ := strconv.Atoi(ins[2])
				ans += m * n
			}
		}
	}

	return ans
}

func ParseInput(filename string) string {
	var a []string

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		a = append(a, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return strings.Join(a, "")
}

func Test(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("Failed! Expected %d to be %d\n", expected, actual))
	} else {
		fmt.Println("Passed!")
	}
}
