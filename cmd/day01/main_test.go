package main

import "testing"

func TestGetCalibrationValue(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{"123456", 16},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tc := range testCases {
		got, err := getCalibrationValue(tc.in)
		if err != nil {
			t.Errorf("getCalibrationValue(%q) returned an unexpected error: %v", tc.in, err)
		} else if got != tc.want {
			t.Errorf("getCalibrationValue(%q) = %d, want %d", tc.in, got, tc.want)
		}
	}
}
