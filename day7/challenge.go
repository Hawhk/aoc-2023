package day7

import (
	"sort"
	"strconv"
	"strings"
)

const (
	HandSize = 5
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	hand     [HandSize]int
	worth    int
	handType HandType
}

func Part1(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")

	types := []HandType{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPairs, OnePair, HighCard}

	handsByType := make(map[HandType][]Hand)

	for _, line := range lines {

		handWorth := strings.Fields(line)
		handStr := handWorth[0]
		worth, err := strconv.Atoi(handWorth[1])

		if err != nil {
			panic("Invalid input")
		}

		hand := getHandAsInts(handStr, 11)
		handType := getHandType(hand, false)
		handsByType[handType] = append(handsByType[handType], Hand{hand, worth, handType})
	}

	for _, hands := range handsByType {
		sort.Slice(hands, func(i, j int) bool {
			for k := 0; k < HandSize; k++ {
				if hands[i].hand[k] > hands[j].hand[k] {
					return true
				} else if hands[i].hand[k] < hands[j].hand[k] {
					return false
				}
			}

			return false
		})
	}

	points := len(lines)

	for _, handType := range types {
		hands := handsByType[handType]

		for _, hand := range hands {
			result += hand.worth * points
			points--
		}

	}

	return result
}

func Part2(input string) int {
	result := 0
	lines := strings.Split(input, "\r\n")
	types := []HandType{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPairs, OnePair, HighCard}

	handsByType := make(map[HandType][]Hand)

	for _, line := range lines {

		handWorth := strings.Fields(line)
		handStr := handWorth[0]
		worth, err := strconv.Atoi(handWorth[1])

		if err != nil {
			panic("Invalid input")
		}

		hand := getHandAsInts(handStr, 1)
		handType := getHandType(hand, true)
		handsByType[handType] = append(handsByType[handType], Hand{hand, worth, handType})
	}

	for _, hands := range handsByType {
		sort.Slice(hands, func(i, j int) bool {
			for k := 0; k < HandSize; k++ {
				if hands[i].hand[k] > hands[j].hand[k] {
					return true
				} else if hands[i].hand[k] < hands[j].hand[k] {
					return false
				}
			}

			return false
		})
	}

	points := len(lines)

	for _, handType := range types {
		hands := handsByType[handType]

		for _, hand := range hands {
			result += hand.worth * points
			points--
		}

	}

	return result
}

func getHandAsInts(handStr string, jValue int) [HandSize]int {

	hand := [HandSize]int{}

	for i, card := range handStr {
		value := 0

		switch string(card) {
		case "T":
			value = 10
		case "J":
			value = jValue
		case "Q":
			value = 12
		case "K":
			value = 13
		case "A":
			value = 14
		default:
			value, _ = strconv.Atoi(string(card))
		}

		hand[i] = value
	}
	return hand
}

func getHandType(hand [HandSize]int, wildCard bool) HandType {

	cards := make(map[int]int)

	for _, card := range hand {
		cards[card]++
	}

	wildCards := cards[1]
	highest := 14

	for card, value := range cards {
		if value > cards[highest] && (card != 1 || !wildCard) {
			highest = card
		}
	}

	if wildCard && wildCards > 0 {
		delete(cards, 1)
		cards[highest] += wildCards
	}

	switch len(cards) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for _, value := range cards {
			if value == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPairs
	case 2:
		if cards[highest] == 1 || cards[highest] == 4 {
			return FourOfAKind
		}
		return FullHouse
	case 1:
		return FiveOfAKind
	default:
		panic("Invalid hand" + strconv.Itoa(len(cards)))
	}

}
