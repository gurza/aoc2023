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

	sum, err := solver.SumAdjacentLines(getGearRatio, 1, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}

type number struct {
	// value represents the numerical value of the number.
	value int

	// startIdx indicates the starting index of the number in the string.
	startIdx int

	// endIdx indicates the ending index of the number in the string.
	endIdx int

	// checked denotes whether the number has been checked or verified.
	checked bool
}

// newNumber converts a string buffer to a number struct, setting its start
// and end indices. It checks if the number is adjacent to a symbol
// in the provided context string 's'.
func newNumber(buf strings.Builder, start, end int, s string) (number, error) {
	val, err := strconv.Atoi(buf.String())
	if err != nil {
		return number{}, fmt.Errorf("failed to parse integer: %w", err)
	}
	chkd := false
	if (start > 0 && isSymbol(rune(s[start-1]))) ||
		(end < len(s)-1 && isSymbol(rune(s[end+1]))) {
		chkd = true
	}

	return number{
		value:    val,
		startIdx: start,
		endIdx:   end,
		checked:  chkd,
	}, nil
}

func sumNumbersAdjacentToSymbols(batch []string, idx int) (int, error) {
	sum := 0

	nums, err := extractNumbers(batch[idx])
	if err != nil {
		return 0, err
	}
	for _, num := range nums {
		if num.checked {
			sum += num.value
			continue
		}

		// Check for symbols in adjacent lines
		for i := range batch {
			if i != idx && hasSymbolsInSubstring(batch[i], num.startIdx-1, num.endIdx+1) {
				sum += num.value
				break
			}
		}
	}

	return sum, nil
}

// isSymbol checks if a character is a symbol, which is defined as being
// neither a digit nor a period.
func isSymbol(ch rune) bool {
	return !unicode.IsDigit(ch) && ch != '.'
}

// hasSymbolsInSubstring checks if there are any symbols in the substring of s
// defined by start and end indices.
func hasSymbolsInSubstring(s string, start, end int) bool {
	if start < 0 {
		start = 0
	}
	if end >= len(s) {
		end = len(s) - 1
	}
	if start > end {
		return false
	}

	for _, ch := range s[start : end+1] {
		if isSymbol(ch) {
			return true
		}
	}

	return false
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
			num, err := newNumber(buf, start, i-1, s)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
			buf.Reset()
			start = -1
		}
	}

	if buf.Len() > 0 {
		num, err := newNumber(buf, start, len(s)-1, s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// extractGears returns the indices of '*' symbols in the input string.
func extractGears(s string) ([]int, error) {
	var idxs []int
	for idx, ch := range s {
		if ch == '*' {
			idxs = append(idxs, idx)
		}
	}
	return idxs, nil
}

func getGearRatio(batch []string, idx int) (int, error) {
	sum := 0

	gears, err := extractGears(batch[idx])
	if err != nil {
		return 0, err
	}

	for _, gear := range gears {
		var agg []number
		for _, s := range batch {
			nums, err := extractNumbers(s)
			if err != nil {
				return 0, err
			}
			for _, num := range nums {
				if gear >= (num.startIdx-1) && gear <= (num.endIdx+1) {
					agg = append(agg, num)
				}
			}
		}
		if len(agg) > 1 {
			m := 1
			for _, num := range agg {
				m *= num.value
			}
			sum += m
		}
	}

	return sum, nil
}
