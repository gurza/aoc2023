package main

import (
	"bufio"
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
	sum, err := solve(fn)
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

func solve(fn string) (int, error) {
	f, err := os.Open(fn)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	sum := 0
	ln := 0
	for scan.Scan() {
		ln++
		v, err := getCalibrationValue(scan.Text())
		if err != nil {
			return 0, fmt.Errorf("error at line %d (%s): %w", ln, scan.Text(), err)
		}
		sum += v
	}

	if err := scan.Err(); err != nil {
		return 0, fmt.Errorf("reading file: %w", err)
	}

	return sum, nil
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
