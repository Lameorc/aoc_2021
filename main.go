package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Lameorc/aoc_2021/day1"
	"github.com/Lameorc/aoc_2021/day2"
)

func readInput(day string) []string {
	filePath := fmt.Sprintf("./inputs/day%s.txt", day)
	// Should fit in memory easily
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

type Solver interface {
	Solve([]string)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	d1 := day1.Day{}
	d1.Solve(readInput("1"))

	d2 := day2.Day2{}
	d2.Solve(readInput("2"))

}
