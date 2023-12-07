package solver

import (
	"fmt"
	"strings"
	"testing"
)

func sumInts(batch []string) (int, error) {
	sum := 0
	for _, s := range batch {
		var v int
		if _, err := fmt.Sscanf(s, "%d", &v); err != nil {
			return 0, err
		}
		sum += v
	}
	return sum, nil
}

func TestSumAdjacentLines(t *testing.T) {
	testCases := []struct {
		h    BatchHandler
		n    int
		in   string
		want int
	}{
		{
			h:  sumInts,
			n:  1,
			in: "0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			// For this test case, each line is added 3 times except the first
			// and last lines, which are added twice
			want: 2*(0+9) + 3*(1+2+3+4+5+6+7+8), // 126
		},
		{
			h:    sumInts,
			n:    2,
			in:   "0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			want: 198,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			rdr := strings.NewReader(tc.in)
			got, err := SumAdjacentLines(tc.h, tc.n, rdr)
			if err != nil {
				t.Errorf("SumAdjacentLines returned an error: %v", err)
			} else if got != tc.want {
				t.Errorf("SumAdjacentLines = %d; want %d", got, tc.want)
			}
		})
	}
}
