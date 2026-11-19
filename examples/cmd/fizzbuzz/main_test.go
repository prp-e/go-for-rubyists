package main

import "testing"

func TestLabel(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := label(tc.input); got != tc.want {
				t.Errorf("label(%d) = %q, want %q",
					tc.input, got, tc.want)
			}
		})
	}
}
