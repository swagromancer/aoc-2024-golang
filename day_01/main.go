package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	// Part 1
	Test(PartOne(ParseInput("input_example.txt")), 11)
	fmt.Println("Part One:", PartOne(ParseInput("input.txt")))

	// Part 2
	Test(PartTwo(ParseInput("input_example.txt")), 31)
	fmt.Println("Part Two:", PartTwo(ParseInput("input.txt")))
}

func PartOne(a, b []int) int {
	slices.Sort(a)
	slices.Sort(b)

	ans := 0

	for i, v := range a {
		ans += AbsInt(v - b[i])
	}

	return ans
}

func PartTwo(a, b []int) int {
	freq := make(map[int]int)
	ans := 0

	for _, n := range b {
		if val, ok := freq[n]; ok {
			freq[n] = val + 1
		} else {
			freq[n] = 1
		}
	}

	for _, n := range a {
		if val, ok := freq[n]; ok {
			ans += n * val
		}
	}

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

func ParseInput(filename string) ([]int, []int) {
	var a, b []int

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		re := regexp.MustCompile(`\s+`)
		line := re.Split(scanner.Text(), -1)

		m, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}

		n, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		a = append(a, m)
		b = append(b, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return a, b
}

func Test(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("Failed! Expected %d to be %d\n", expected, actual))
	} else {
		fmt.Println("Passed!")
	}
}
