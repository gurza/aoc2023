package main

import (
	"fmt"
	"testing"
)

func TestCheckCubesAndReturnGameID(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			1,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			2,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			0,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			0,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			5,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got, err := checkCubesAndReturnGameID(tc.in)
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.in, err)
			} else if got != tc.want {
				t.Errorf("checkCubesAndReturnGameID(%s) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

func TestGetSetPower(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			48,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			12,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			1560,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			630,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			36,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got, err := getSetPower(tc.in)
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.in, err)
			} else if got != tc.want {
				t.Errorf("getSetPower(%s) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
