package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct {}


type pos struct {
	depth int
	horizontal int
}

type part2Pos struct {
	depth, horizontal, aim int
}

type moveFun func (p *pos, amount int)

type moveFunP2 func (p *part2Pos, amount int)

var directionFunMap = map[string]moveFun{
	"forward": func(p *pos, amount int) {p.horizontal += amount},
	"down": func(p *pos, amount int) {p.depth += amount},
	"up": func(p *pos, amount int) {p.depth -= amount},
}

var directionFunMapP2 = map[string]moveFunP2{
	"forward": func(p *part2Pos, amount int) {
		p.horizontal += amount
		p.depth += p.aim * amount
	},
	"down": func(p *part2Pos, amount int) {p.aim += amount},
	"up": func(p *part2Pos, amount int) {p.aim -= amount},
}



func (d *Day2) Solve (input []string) {
	p := &pos{}
	p2 := &part2Pos{}
	for _, line := range input {
		if line == "" {
			continue
		}
		splits := strings.Split(line, " ")
		direction := splits[0]
		amount, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Printf("Failed to convert %s: %v", line, err)
			continue
		}
		directionFunMap[direction](p, amount)
		directionFunMapP2[direction](p2, amount)
	}
	fmt.Printf("Day 2, Part 1: %d\n", p.depth * p.horizontal)
	fmt.Printf("Day 2, Part 2: %d\n", p2.depth * p2.horizontal)
}

