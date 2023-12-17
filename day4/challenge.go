package day4

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")

	for _, line := range lines {
		points := 0
		winningMap := make(map[string]bool)
		game := strings.Split(line, ":")
		sections := strings.Split(game[1], "|")

		winings := strings.Split(sections[0], " ")
		nums := strings.Split(sections[1], " ")

		for _, win := range winings {
			winningMap[win] = true
		}
		for _, num := range nums {
			if _, ok := winningMap[num]; ok {
				if _, err := strconv.Atoi(num); err == nil {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
				}
			}
		}
		result += points

	}

	return result
}

func Part2(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")

	cards := []int{}

	for a := 0; a < len(lines); a++ {
		cards = append(cards, a)
	}

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		points := 0
		winningMap := make(map[string]bool)

		game := strings.Split(lines[card], ":")
		sections := strings.Split(game[1], "|")

		winings := strings.Split(sections[0], " ")
		nums := strings.Split(sections[1], " ")

		for _, win := range winings {
			winningMap[win] = true
		}

		for _, num := range nums {
			if _, ok := winningMap[num]; ok {
				if _, err := strconv.Atoi(num); err == nil {
					points += 1
					cards = append(cards, card+points)
				}
			}
		}
	}

	result = len(cards)
	return result
}
