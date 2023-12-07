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
func SumAdjacentLines(h BatchHandler, n int, r io.Reader) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("n must be at least 1")
	}

	scan := bufio.NewScanner(r)
	var window []string
	sum := 0
	i := 0

	processBatch := func() error {
		v, err := h(window)
		if err != nil {
			return fmt.Errorf("batch error at line %d: %w", i, err)
		}
		sum += v
		return nil
	}

	for scan.Scan() {
		line := scan.Text()

		if len(window) == 2*n+1 {
			window = window[1:]
		}
		window = append(window, line)

		if i >= n {
			if err := processBatch(); err != nil {
				return 0, err
			}
		}

		i++
	}

	if err := scan.Err(); err != nil {
		return 0, fmt.Errorf("failed to read from input: %w", err)
	}

	for j := 0; j < n; j++ {
		window = window[1:]
		if err := processBatch(); err != nil {
			return 0, err
		}
	}

	return sum, nil
}
