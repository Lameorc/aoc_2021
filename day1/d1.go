package day1

import (
	"fmt"
	"strconv"
)

type Day struct{}

func (d *Day) Solve(input []string) {
	result := 0
	resultPart2 := 0
	prev := 99999
	inputLen := len(input)
	for i, l := range input {
		if l == "" {
			continue
		}
		currentWindowSum := 0
		nextWindowSum := 0
		for j := 0; j < 3 && i+j < inputLen; j++ {
			val, err := strconv.ParseInt(input[i + j], 0, 32)
			if err != nil {
				continue
			}
			aVal := int(val)
			currentWindowSum += aVal
		}
		for j := 1; j < 4 && i+j < inputLen; j++ {
			val, err := strconv.ParseInt(input[i + j], 0, 32)
			if err != nil {
				continue
			}
			aVal := int(val)
			nextWindowSum += aVal
		}
		if nextWindowSum > currentWindowSum {
			resultPart2++
		}
		lval, err := strconv.ParseInt(l, 0, 32)
		if err != nil {
			fmt.Print(err)
			return
		}

		newVal := int(lval)
		if newVal > prev {
			result += 1
		}
		prev = newVal

	}
	fmt.Printf("Day 1, part 1: %d\n", result)
	fmt.Printf("Day 1, part 2: %d\n", resultPart2)
}
