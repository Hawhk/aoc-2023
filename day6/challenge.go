package day6

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	result := 1
	lines := strings.Split(input, "\r\n")

	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])

	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		if err != nil {
			continue
		}

		won := 0

		for j := 1; j < time; j++ {

			dist := j * (time - j)

			if dist > distance {
				won++
			}
		}

		result *= won
	}

	return result
}

func Part2(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")

	time, err1 := strconv.Atoi(strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), ""))
	distance, err2 := strconv.Atoi(strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), ""))

	if err1 != nil || err2 != nil {
		panic("Invalid input")
	}

	for j := 1; j < time; j++ {

		dist := j * (time - j)

		if dist > distance {
			result++
		}
	}

	return result
}
