package solver

import (
	"bufio"
	"fmt"
	"io"
)

// Handler defines the type for the function that will be called
// on each input string.
type Handler func(s string) (int, error)

// SumLines reads lines from the given reader, processes each line
// with the provided handler, and returns the sum.
func SumLines(h Handler, r io.Reader) (int, error) {
	sum := 0

	scan := bufio.NewScanner(r)
	i := 0
	for scan.Scan() {
		v, err := h(scan.Text())
		if err != nil {
			return 0, fmt.Errorf("line %d: %v, %w", i+1, scan.Text(), err)
		}
		sum += v
		i++
	}

	if err := scan.Err(); err != nil {
		return 0, fmt.Errorf("failed to read from input: %w", err)
	}

	return sum, nil
}

// BatchHandler is a function type that takes a batch of strings and an index.
// 'batch' is a slice of strings, where each string is subject to processing.
// 'idx' is the index of the string currently being processed in the batch.
type BatchHandler func(batch []string, idx int) (int, error)

// SumAdjacentLines reads lines from the given reader, processes batches of
// lines, including each line and its adjacent lines, with the given handler.
// It returns the sum of the processed values.
func SumAdjacentLines(h BatchHandler, n int, r io.Reader) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("n must be at least 1")
	}

	var buf []string
	sum := 0

	proc := func(i int) error {
		v, err := h(buf, i)
		if err != nil {
			return fmt.Errorf("batch error: %w", err)
		}
		sum += v
		return nil
	}

	scan := bufio.NewScanner(r)
	i := 0
	for scan.Scan() {
		line := scan.Text()

		if len(buf) == 2*n+1 {
			buf = buf[1:]
		}
		buf = append(buf, line)

		if i >= n {
			if err := proc(i - n); err != nil {
				return 0, err
			}
		}

		i++
	}

	if err := scan.Err(); err != nil {
		return 0, fmt.Errorf("failed to read from input: %w", err)
	}

	for j := 0; j < n; j++ {
		if len(buf) == 0 {
			break
		}

		buf = buf[1:]
		if err := proc(i - j); err != nil {
			return 0, err
		}
	}

	return sum, nil
}
