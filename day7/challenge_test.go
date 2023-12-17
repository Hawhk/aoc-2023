package day7

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	data := readData(1)
	whant := 6440

	got := Part1(data)

	if got != whant {
		t.Errorf("got %d, want %d", got, whant)
	}
}

func TestPart2(t *testing.T) {
	data := readData(2)
	whant := 5905

	got := Part2(data)

	if got != whant {
		t.Errorf("got %d, want %d", got, whant)
	}
}

func readData(part int) string {
	content, err := os.ReadFile(fmt.Sprintf("./data/test-%d.txt", part))
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
