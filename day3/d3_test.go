package day3

import "testing"

var input = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1(t *testing.T) {
	result := part1(input)
	if result.epsilonRate != 9 {
		t.Fatalf("Wrong epsilon rate, expected 9, got %d(%b)", result.epsilonRate, result.epsilonRate)
	}
	if result.gammaRate != 22 {
		t.Fatalf("Wrong gamma rate, expected 22, got %d(%b)", result.gammaRate, result.gammaRate)
	}

}

func TestPart2(t *testing.T) {

}
