package day3

import (
	"fmt"
	"strconv"
)


type Day struct {}

type Result struct{
	gammaRate, epsilonRate, oxy, co2 int
}

func part1(input []string) Result {

	occurrences := make(map[int]map[int]int)
	for _, line := range input {
		if line == "" {
			continue
		}

		for i, r := range line {
			if occurrences[i] == nil{
				occurrences[i] = make(map[int]int)
			}

			val, err := strconv.Atoi(string(r))
			if err != nil {
				continue
			}
			occurrences[i][val]++
		}
	}
	result := Result{}
	for i := range occurrences {
		realIndex := uint(len(occurrences)) - uint(i) -1
		if occurrences[i][0] > occurrences[i][1] {
			result.epsilonRate = setBit(result.epsilonRate, realIndex)
		} else{
			result.gammaRate = setBit(result.gammaRate, realIndex)
		}
	}

	return result
}

func part2(input []string) Result {
	// make a copy first, so we can modify it
	oxyCandidates := make([]int, len(input))
	co2Candidates := make([]int, len(input))
	for i, line := range input {
		if asInt, err := strconv.Atoi(line); err != nil {
			oxyCandidates[i] = asInt
			co2Candidates[i] = asInt
		}

	}
	// guess we should not need more than len attempts
	for i:= 0; i < len(input[0]); i++{
		setOxyOccurrences, unsetOxyOccurences := 0, 0
		// setCO2Occurrences, unsetCO2Occurences := 0, 0

		for _, value := range oxyCandidates {
			if isSet(value, i) {
				setOxyOccurrences++
			} else {
				unsetOxyOccurences++
			}
		}

		oxyValsToKeep := make([]int, 0)

		var keepFunc func (int) bool = nil
		if unsetOxyOccurences > setOxyOccurrences {
			keepFunc = func (val int) bool {
				return !isSet(val, i)
			}
		} else {
			keepFunc = func (val int) bool {
				return isSet(val, i)
			}
		}
		for _, val := range oxyCandidates {
			if !keepFunc(val) {
				oxyValsToKeep = append(oxyValsToKeep, val)
			}
		}
		oxyCandidates := make([]int, 0)
		for _, val := range oxyValsToKeep {
			oxyCandidates = append(oxyCandidates, val)
		}


		// might be able to end early, but I don't think that's likely
		if len(oxyCandidates) == 1 && len(co2Candidates) == 1 {
			break
		}
	}

	oxy := oxyCandidates[0]
	co2 := co2Candidates[0]


	return Result{oxy: oxy, co2: co2}
}


func (d *Day) Solve(input []string)  {
	res1 := part1(input)
	res2 := part2(input)
	fmt.Printf("Day 3, part 1: %d\n", res1.epsilonRate * res1.gammaRate)
	fmt.Printf("Day 3, part 2: %d\n", res2.oxy * res2.co2)
}

// Checks wheter bit at pos in integer n is set.
func isSet(n, pos int) bool {
	n2 := (n >> (pos - 1 ))

	return (n2 & 1) == 1

}

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}
