package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumNumbersAdjacentToSymbols(t *testing.T) {
	testCases := []struct {
		batch []string
		idx   int
		want  int
	}{
		{
			batch: []string{
				"467..114..",
				"...*......",
			},
			idx:  0,
			want: 467,
		},
		{
			batch: []string{
				"467..114..",
				"...*......",
				"..35..633.",
			},
			idx:  1,
			want: 0,
		},
		{
			batch: []string{
				"...*......",
				"..35..633.",
				"......#...",
			},
			idx:  1,
			want: 35 + 633,
		},
		{
			batch: []string{
				"..35..633.",
				"......#...",
				"617*......",
			},
			idx:  1,
			want: 0,
		},
		{
			batch: []string{
				"......#...",
				"617*......",
				".....+.58.",
			},
			idx:  1,
			want: 617,
		},
		{
			batch: []string{
				"617*......",
				".....+.58.",
				"..592.....",
			},
			idx:  1,
			want: 0,
		},
		{
			batch: []string{
				".....+.58.",
				"..592.....",
				"......755.",
			},
			idx:  1,
			want: 592,
		},
		{
			batch: []string{
				"..592.....",
				"......755.",
				"...$.*....",
			},
			idx:  1,
			want: 755,
		},
		{
			batch: []string{
				"......755.",
				"...$.*....",
				".664.598..",
			},
			idx:  1,
			want: 0,
		},
		{
			batch: []string{
				"...$.*....",
				".664.598..",
			},
			idx:  1,
			want: 664 + 598,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got, err := sumNumbersAdjacentToSymbols(tc.batch, tc.idx)
			if err != nil {
				t.Errorf("TestSumNumbersAdjacentToSymbols(%s) error = %v", tc.batch, err)
			}
			if got != tc.want {
				t.Errorf("TestSumNumbersAdjacentToSymbols(%s) = %d; want %d", tc.batch, got, tc.want)
			}
		})
	}
}

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
				{664, 1, 3, false},
				{598, 5, 7, false},
			},
		},
		{
			name: "Numbers at edges",
			in:   "1......755",
			want: []number{
				{1, 0, 0, false},
				{755, 7, 9, false},
			},
		},
		{
			name: "Number adjacent to symbol",
			in:   "617*......",
			want: []number{
				{617, 0, 2, true},
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
