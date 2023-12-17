package day6

import (
	"log"
	"os"
	"testing"
)

var data = readData()

func TestPart1(t *testing.T) {
	whant := 288

	got := Part1(data)

	if got != whant {
		t.Errorf("got %d, want %d", got, whant)
	}
}

func TestPart2(t *testing.T) {
	whant := 71503

	got := Part2(data)

	if got != whant {
		t.Errorf("got %d, want %d", got, whant)
	}
}

func readData() string {
	content, err := os.ReadFile("./data/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
