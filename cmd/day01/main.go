package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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
	var first, last rune

	found := false
	for _, r := range s {
		if unicode.IsDigit(r) {
			if !found {
				first = r
				found = true
			}
			last = r
		}
	}

	if !found {
		return 0, fmt.Errorf("no digits found in input")
	}

	res := string(first) + string(last)
	return strconv.Atoi(res)
}
