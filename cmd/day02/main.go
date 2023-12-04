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

	sum, err := solver.SumLines(f, getSetPower)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}

func getSetPower(s string) (int, error) {
	sets, err := getSets(s)
	if err != nil {
		return 0, err
	}

	red, green, blue := 0, 0, 0
	for _, set := range sets {
		red1, green1, blue1, err := getCubes(set)
		if err != nil {
			return 0, err
		}
		if red1 > red {
			red = red1
		}
		if green1 > green {
			green = green1
		}
		if blue1 > blue {
			blue = blue1
		}
	}

	return red * green * blue, nil
}

func checkCubesAndReturnGameID(s string) (int, error) {
	id, err := getGameID(s)
	if err != nil {
		return 0, err
	}

	sets, err := getSets(s)
	if err != nil {
		return 0, err
	}

	for _, set := range sets {
		red, green, blue, err := getCubes(set)
		if err != nil {
			return 0, err
		}
		if red > 12 || green > 13 || blue > 14 {
			return 0, nil
		}
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

func getSets(gameEntry string) ([]string, error) {
	parts := strings.SplitN(gameEntry, ":", 2)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid game entry format")
	}

	sets := strings.Split(strings.TrimSpace(parts[1]), ";")

	for i, set := range sets {
		sets[i] = strings.TrimSpace(set)
	}

	return sets, nil
}

func getCubes(s string) (int, int, int, error) {
	var red, green, blue int
	var err error

	parts := strings.Split(s, ",")
	for _, part := range parts {
		piece := strings.Fields(strings.TrimSpace(part))
		if len(piece) != 2 {
			return 0, 0, 0, fmt.Errorf("invalid format")
		}

		count, err := strconv.Atoi(piece[0])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid number format: %w", err)
		}

		switch piece[1] {
		case "blue":
			blue += count
		case "red":
			red += count
		case "green":
			green += count
		default:
			return 0, 0, 0, fmt.Errorf("invalid color: %s", piece[1])
		}
	}

	return red, green, blue, err
}
