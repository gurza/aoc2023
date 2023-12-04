package main

import (
	"aoc2023/pkg/solver"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// digits are actually spelled out with letters
var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	fn := parseFlags()
	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", fn, err)
		os.Exit(1)
	}
	defer f.Close()

	sum, err := solver.SumLines(f, getCalibrationValue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}

func parseFlags() string {
	fn := flag.String("f", "", "Path to your puzzle input file")
	flag.Parse()

	if *fn == "" {
		fmt.Println("Usage: program -f <filename>")
		os.Exit(1)
	}

	return *fn
}

func getCalibrationValue(s string) (int, error) {
	var first, last int

	for i := 0; i < len(s); i++ {
		dig, ok := findDigit(s[i:])
		if ok {
			if first == 0 {
				first = dig
			}
			last = dig
		}
	}

	if first == 0 {
		return 0, fmt.Errorf("no digits found in input")
	}

	return first*10 + last, nil
}

func findDigit(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	if dig, ok := checkNumericDigit(rune(s[0])); ok {
		return dig, true
	}

	if dig, ok := checkSpelledDigit(s); ok {
		return dig, true
	}

	return 0, false
}

func checkNumericDigit(r rune) (int, bool) {
	if unicode.IsDigit(r) {
		dig, _ := strconv.Atoi(string(r))
		return dig, true
	}
	return 0, false
}

func checkSpelledDigit(s string) (int, bool) {
	for word, dig := range digits {
		if len(s) >= len(word) && s[:len(word)] == word {
			return dig, true
		}
	}
	return 0, false
}
