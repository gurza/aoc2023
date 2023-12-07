package main

import (
	"reflect"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want []number
	}{
		{
			name: "No numbers",
			in:   "...*......",
			want: nil,
		},
		{
			name: "Multiple digit numbers",
			in:   ".664.598..",
			want: []number{
				{664, 1, 3},
				{598, 5, 7},
			},
		},
		{
			name: "Numbers at edges",
			in:   "1......755",
			want: []number{
				{1, 0, 0},
				{755, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := extractNumbers(tc.in)
			if err != nil {
				t.Errorf("TestExtractNumbers(%s) error = %v", tc.name, err)
			} else if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TestExtractNumbers(%s) = %v; want %v", tc.name, got, tc.want)
			}
		})
	}
}
