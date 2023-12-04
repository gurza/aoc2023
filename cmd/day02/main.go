package main

import (
	"aoc2023/pkg/solver"
	"aoc2023/pkg/util"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	id, err := getGameID(s)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getGameID(s string) (int, error) {
	parts := strings.Split(s, ":")
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid game entry format")
	}

	id := strings.TrimSpace(strings.TrimPrefix(parts[0], "Game"))

	id1, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("invalid game ID: %w", err)
	}

	return id1, nil
}
