package main

import "testing"

func TestLeastInterval(t *testing.T) {
	cases := []struct {
		tasks string
		n     int
		want  int
	}{
		{"AAABBB", 2, 8}, // A B _ A B _ A B
		{"AAABBB", 0, 6},
		{"AAAAAABCDEFG", 2, 16},
		{"AAAB", 2, 7}, // A _ _ A _ _ A then B inside: A B _ A _ _ A = 7
		{"ABCDEF", 2, 6},
		{"A", 5, 1},
		{"", 3, 0},
		{"AAAABBBBCCD", 2, 11}, // ties on max frequency fill the final row
	}
	for _, c := range cases {
		if got := LeastInterval([]byte(c.tasks), c.n); got != c.want {
			t.Errorf("LeastInterval(%q,%d) = %d, want %d", c.tasks, c.n, got, c.want)
		}
	}
}
