package main

import (
	"aoc2023/pkg/solver"
	"aoc2023/pkg/util"
	"fmt"
	"os"
)

func main() {
	fn, err := util.ParseFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", fn, err)
		os.Exit(1)
	}
	defer f.Close()

	sum, err := solver.SumLines(f, getSumIDs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}

func getSumIDs(s string) (int, error) {
	return 1, nil
}
