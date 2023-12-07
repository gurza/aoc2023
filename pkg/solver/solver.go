package solver

import (
	"bufio"
	"fmt"
	"io"
)

// Handler defines the type for the function that will be called
// on each input string.
type Handler func(string) (int, error)

// SumLines reads lines from the given reader, processes each line
// with the provided handler, and returns the sum.
func SumLines(r io.Reader, h Handler) (int, error) {
	sum := 0

	scan := bufio.NewScanner(r)
	i := 0
	for scan.Scan() {
		i++
		v, err := h(scan.Text())
		if err != nil {
			return 0, fmt.Errorf("line %d: %v, %w", i, scan.Text(), err)
		}
		sum += v
	}

	if err := scan.Err(); err != nil {
		return 0, fmt.Errorf("failed to read from input: %w", err)
	}

	return sum, nil
}

type BatchHandler func([]string) (int, error)

// SumAdjacentLines reads lines from the given reader, processes batches of
// lines, including each line and its adjacent lines, with the given handler.
// It returns the sum of the processed values.
func SumAdjacentLines(r io.Reader, n int, h BatchHandler) (int, error) {
	return 0, nil
}
