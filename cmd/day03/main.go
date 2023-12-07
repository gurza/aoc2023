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
	// value represents the numerical value of the number.
	value int

	// startIdx indicates the starting index of the number in the string.
	startIdx int

	// endIdx indicates the ending index of the number in the string.
	endIdx int

	// checked denotes whether the number has been checked or verified.
	checked bool
}

func sumAdjacentNumbers(b []string) (int, error) {
	sum := 0
	return sum, nil
}

// isSymbol checks if a character is a symbol, which is defined as being
// neither a digit nor a period.
func isSymbol(ch rune) bool {
	return !unicode.IsDigit(ch) && ch != '.'
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
			chkd := false
			if start > 0 && isSymbol(rune(s[start-1])) {
				chkd = true
			}
			if i < len(s) && isSymbol(rune(s[i])) {
				chkd = true
			}
			nums = append(nums, number{
				value:    val,
				startIdx: start,
				endIdx:   i - 1,
				checked:  chkd,
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
		chkd := false
		if start > 0 && isSymbol(rune(s[start-1])) {
			chkd = true
		}
		if len(s) > start+buf.Len() && isSymbol(rune(s[start+buf.Len()])) {
			chkd = true
		}
		nums = append(nums, number{
			value:    value,
			startIdx: start,
			endIdx:   len(s) - 1,
			checked:  chkd,
		})
	}

	return nums, nil
}
