package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gitlab.com/Newmark/aoc_2021/day1"
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

}