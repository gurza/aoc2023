package main

import (
	"aoc2023/pkg/solver"
	"aoc2023/pkg/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

	sum, err := solver.SumAdjacentLines(sumAdjacentNumbers, 1, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}

type number struct {
	value    int
	startIdx int
	endIdx   int
}

func sumAdjacentNumbers(b []string) (int, error) {
	sum := 0
	return sum, nil
}

func extractNumbers(s string) ([]number, error) {
	var nums []number

	var buf strings.Builder
	start := -1
	for i, ch := range s {
		if unicode.IsDigit(ch) {
			if start == -1 {
				start = i
			}
			buf.WriteRune(ch)
		} else if buf.Len() > 0 {
			val, err := strconv.Atoi(buf.String())
			if err != nil {
				return nil, err
			}
			nums = append(nums, number{
				value:    val,
				startIdx: start,
				endIdx:   i - 1,
			})
			buf.Reset()
			start = -1
		}
	}

	if buf.Len() > 0 {
		value, err := strconv.Atoi(buf.String())
		if err != nil {
			return nil, err
		}
		nums = append(nums, number{
			value:    value,
			startIdx: start,
			endIdx:   len(s) - 1,
		})
	}

	return nums, nil
}
