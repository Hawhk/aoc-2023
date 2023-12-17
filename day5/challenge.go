package day5

import (
	"math"
	"strconv"
	"strings"
	"sync"
)

type Mapper struct {
	Destination int
	Source      int
	Range       int
}

func Part1(input string) int {
	result := math.MaxInt64
	lines := strings.Split(input, "\r\n")

	a := strings.Split(lines[0], ":")
	seeds := strings.Split(a[1], " ")

	mappers := map[int][]Mapper{}

	mapNum := 0

	for i := 3; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		if strings.Contains(line, "map:") {
			mapNum++
			continue
		}

		numStrs := strings.Split(line, " ")
		nums := []int{}

		for _, numStr := range numStrs {
			if num, err := strconv.Atoi(numStr); err == nil {
				nums = append(nums, num)
			}
		}

		mappers[mapNum] = append(mappers[mapNum], Mapper{
			Destination: nums[0],
			Source:      nums[1],
			Range:       nums[2],
		})
	}

	for _, seed := range seeds {
		var value int

		if num, err := strconv.Atoi(seed); err == nil {
			value = num
		} else {
			continue
		}

		result = min(ab(mappers, value), result)
	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\r\n")
	a := strings.Split(lines[0], ":")
	seeds := strings.Split(a[1], " ")

	mappers := map[int][]Mapper{}

	mapNum := 0
	for i := 3; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		if strings.Contains(line, "map:") {
			mapNum++
			continue
		}

		numStrs := strings.Split(line, " ")
		nums := []int{}

		for _, numStr := range numStrs {
			if num, err := strconv.Atoi(numStr); err == nil {
				nums = append(nums, num)
			}
		}

		mappers[mapNum] = append(mappers[mapNum], Mapper{
			Destination: nums[0],
			Source:      nums[1],
			Range:       nums[2],
		})
	}

	seedNums := []int{}

	for _, seed := range seeds {
		if num, err := strconv.Atoi(seed); err == nil {
			seedNums = append(seedNums, num)
		}
	}

	var wg sync.WaitGroup
	results := make(chan int)
	result := make(chan int)

	for i := 0; i < len(seedNums); i += 2 {
		wg.Add(1)
		go abGo(mappers, i, seedNums, &wg, results)
	}

	go collect(results, result)

	wg.Wait()
	close(results)

	return <-result
}

func abGo(mappers map[int][]Mapper, i int, seedNums []int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	minVal := math.MaxInt64
	for j := seedNums[i]; j < seedNums[i]+seedNums[i+1]; j++ {
		minVal = min(minVal, ab(mappers, j))
	}
	results <- minVal
}

func collect(results <-chan int, result chan<- int) {

	minVal := math.MaxInt64
	for result := range results {
		minVal = min(minVal, result)
	}
	result <- minVal
}

func ab(mappers map[int][]Mapper, value int) int {
	for i := 0; i < len(mappers); i++ {
		maps := mappers[i]
		for _, map_ := range maps {
			if value >= map_.Source && value < map_.Source+map_.Range {
				value = map_.Destination + (value - map_.Source)
				break
			}
		}
	}
	return value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
