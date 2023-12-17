package day3

import (
	"strings"
	"unicode"
)

func Part1(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")

	for row, line := range lines {
		number, numIndex, numLen := 0, -1, 0
		for col, char := range line {
			if unicode.IsDigit(char) {
				if numIndex == -1 {
					numIndex = col
				} else {
					number *= 10
				}
				numLen++
				number += int(char - '0')
			}
			if (numIndex != -1 && !unicode.IsDigit(char)) || col == len(line)-1 {
				for i := -1; i <= 1; i++ {
					if row+i < 0 || row+i >= len(lines) {
						continue
					}
					runes := []rune(lines[row+i])
					for j := numIndex - 1; j <= numIndex+numLen; j++ {
						if j < 0 || j >= len(runes) || (i == 0 && j >= numIndex && j < numIndex+numLen) {
							continue
						}
						rune := runes[j]
						// fmt.Print(row+i, j, string(rune), " ")
						if !unicode.IsDigit(rune) && rune != '.' {
							result += number
							// fmt.Println("found!")
							break
						}
					}
				}
				numIndex, numLen, number = -1, 0, 0
			}
		}
	}

	return result
}

func Part2(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")

	for row, line := range lines {
		for col, char := range line {
			if char == '*' {
				touching := 0
				ratio := 1
				for i := -1; i <= 1; i++ {
					nums := getNums(lines[row+i])
					for _, num := range nums {
						if num[1]-1 <= col && num[2]+1 >= col {
							touching++
							ratio *= num[0]
						}
					}
				}
				if touching == 2 {
					result += ratio
				}
			}
		}
	}

	return result
}

func getNums(String string) [][]int {
	var nums [][]int
	sep := true
	for col, char := range String {

		length := len(nums) - 1
		if unicode.IsDigit(char) {
			x := int(char - '0')
			if sep {
				nums = append(nums, []int{x, col, col})
				sep = false
			} else {
				nums[length][0] = nums[length][0]*10 + x
				nums[length][2] = nums[length][2] + 1
			}
		} else {
			sep = true
		}
	}
	return nums
}
