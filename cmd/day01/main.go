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
	if err := run(fn); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
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

func run(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	return nil
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
