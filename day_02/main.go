package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Part 1
	Test(PartOne(ParseInput("input_example.txt")), 2)
	fmt.Println("Part One:", PartOne(ParseInput("input.txt")))

	// Part 2
	// Test(PartTwo(ParseInput("input_example.txt")), 31)
	// fmt.Println("Part Two:", PartTwo(ParseInput("input.txt")))
}

func PartOne(input [][]int) int {
	ans := 0

	for _, report := range input {
		if ValidateReport(report) {
			ans++
		}
	}
	return ans
}

func ValidateReport(r []int) bool {
	// A report passes if:
	// The levels are either all increasing or all decreasing.
	// Any two adjacent levels differ by at least one and at most three.
	isInc, isDec := false, false
	for i := 1; i < len(r); i++ {
		if r[i] > r[i-1] {
			isInc = true
		} else {
			isDec = true
		}

		if isInc && isDec {
			return false
		}

		diff := AbsInt(r[i] - r[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func PartTwo(input [][]int) int {
	ans := 0

	return ans
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func ParseInput(filename string) [][]int {
	var a [][]int

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		re := regexp.MustCompile(`\s`)
		line := re.Split(scanner.Text(), -1)
		ints := make([]int, len(line))

		for i, s := range line {
			ints[i], _ = strconv.Atoi(s)
		}

		a = append(a, ints)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return a
}

func Test(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("Failed! Expected %d to be %d\n", expected, actual))
	} else {
		fmt.Println("Passed!")
	}
}
