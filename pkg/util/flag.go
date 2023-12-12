package util

import (
	"flag"
	"fmt"
)

// ParseFlags parses command line flags and returns the filename or an error.
func ParseFlags() (string, error) {
	fn := flag.String("f", "", "Path to your puzzle input file")
	flag.Parse()

	if *fn == "" {
		return "", fmt.Errorf("missing file name: usage: program -f <filename>")
	}

	return *fn, nil
}
