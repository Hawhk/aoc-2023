package day2

import (
	"strconv"
	"strings"
)

var colors = []string{"red", "green", "blue"}

func Part1(input string) int {
	result := 0
	lines := strings.Split(input, "\n")

	colorsCompare := []int{12, 13, 14}

	for _, line := range lines {
		colorsPicked := []int{0, 0, 0}

		words := strings.Fields(line)
		id := 0
		for wordIndex, word := range words {
			if val, err := strconv.Atoi(word); err == nil {
				for i, color := range colors {
					if strings.Contains(words[wordIndex+1], color) {
						colorsPicked[i] = val
						break
					}
				}
			}

			if strings.Contains(word, ";") || wordIndex == len(words)-1 {
				for i, colorCompare := range colorsCompare {
					if colorCompare < colorsPicked[i] {
						id = 0
						break
					}
				}
				colorsPicked = []int{0, 0, 0}
			} else if strings.Contains(word, ":") {
				if val, err := strconv.Atoi(word[0 : len(word)-1]); err == nil {
					id = val
				}
			}
		}
		result += id
	}

	return result
}

func Part2(input string) int {
	result := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		colorsPicked := []int{0, 0, 0}
		words := strings.Fields(line)
		for wordIndex, word := range words {
			if val, err := strconv.Atoi(word); err == nil {
				for i, color := range colors {
					if strings.Contains(words[wordIndex+1], color) {
						colorsPicked[i] = max(colorsPicked[i], val)
						break
					}
				}
			}
		}

		result += colorsPicked[0] * colorsPicked[1] * colorsPicked[2]
	}

	return result
}
