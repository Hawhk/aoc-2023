package day1

import (
	"strings"
	"unicode"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Part1(input string) int {

	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		firstNumber, lastNumber := -1, -1
		for _, char := range line {
			if unicode.IsDigit(char) { // Try converting word to int

				if firstNumber == -1 {
					firstNumber = int(char-'0') * 10
				}
				lastNumber = int(char - '0')
			}
		}

		result += firstNumber + lastNumber

	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	result := 0

	for _, line := range lines {
		result += convertWordToNumber(line)
	}

	return result
}

func convertWordToNumber(line string) int {
	firstNumber, lastNumber := -1, -1
	for i, char := range line {
		found := -1
		if unicode.IsDigit(char) {
			found = int(char - '0')
		} else {
			found = tryConvertWordToNumber(line[i:])
		}
		if found != -1 {
			if firstNumber == -1 {
				firstNumber = found * 10
			}
			lastNumber = found
		}
	}
	return firstNumber + lastNumber
}

func tryConvertWordToNumber(line string) int {
	found := -1
	for j, num := range numbers {

		last := min(len(num), len(line))

		if strings.Contains(line[:last], num) {
			found = j + 1
			break
		}
	}

	return found
}
