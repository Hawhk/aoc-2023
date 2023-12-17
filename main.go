package main

import (
	"day1"
	"day2"
	"day3"
	"day4"
	"day5"
	"day6"
	"day7"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"sync"
	"time"
)

type Days func(string) int

var days = map[string]Days{
	"day-1-1": day1.Part1,
	"day-1-2": day1.Part2,
	"day-2-1": day2.Part1,
	"day-2-2": day2.Part2,
	"day-3-1": day3.Part1,
	"day-3-2": day3.Part2,
	"day-4-1": day4.Part1,
	"day-4-2": day4.Part2,
	"day-5-1": day5.Part1,
	"day-5-2": day5.Part2,
	"day-6-1": day6.Part1,
	"day-6-2": day6.Part2,
	"day-7-1": day7.Part1,
	"day-7-2": day7.Part2,
}

func main() {
	day, part := "", ""
	if len(os.Args) < 2 && len(os.Args) != 1 {
		fmt.Println("Please enter a day and part number")
		return
	} else if len(os.Args) == 2 {
		day = os.Args[1]
	} else if len(os.Args) == 3 {
		day, part = os.Args[1], os.Args[2]
	}

	if day == "" && part == "" {
		var wg sync.WaitGroup
		var keys []string
		for key := range days {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {

			vals := regexp.MustCompile(`\d+`).FindAllString(key, -1)

			if vals[0] == "0" {
				continue
			}
			wg.Add(1)
			go runChallengeAsync(vals[0], vals[1], &wg)
		}

		wg.Wait()
	} else if day != "" && part == "" {
		runChallenge(day, "1")
		runChallenge(day, "2")
	} else {
		runChallenge(day, part)
	}

}

func runChallengeAsync(day string, part string, wg *sync.WaitGroup) {
	defer wg.Done()
	runChallenge(day, part)
}

func runChallenge(day string, part string) {
	content, err := os.ReadFile(fmt.Sprintf("./day%s/data/%s.txt", day, part))
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now()
	result := days[fmt.Sprintf("day-%s-%s", day, part)](string(content))
	elapsedTime := time.Since(startTime)
	fmt.Printf("Runing day %s part %s\n", day, part)
	fmt.Printf("Result: %d in %s\n", result, elapsedTime)

}
