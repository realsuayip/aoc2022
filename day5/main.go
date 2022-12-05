package day5

import (
	"fmt"
	"strings"
)

var input = "    [G]         [P]         [M]    \n    [V]     [M] [W] [S]     [Q]    \n    [N]     [N] [G] [H]     [T] [F]\n    [J]     [W] [V] [Q] [W] [F] [P]\n[C] [H]     [T] [T] [G] [B] [Z] [B]\n[S] [W] [S] [L] [F] [B] [P] [C] [H]\n[G] [M] [Q] [S] [Z] [T] [J] [D] [S]\n[B] [T] [M] [B] [J] [C] [T] [G] [N]\n 1   2   3   4   5   6   7   8   9 \n\nmove 2 from 4 to 2\nmove 6 from 9 to 7\nmove 4 from 7 to 2\nmove 2 from 4 to 1\nmove 2 from 6 to 7\nmove 1 from 3 to 8\nmove 4 from 7 to 1\nmove 2 from 3 to 2\nmove 3 from 8 to 5\nmove 3 from 1 to 4\nmove 12 from 2 to 5\nmove 2 from 6 to 8\nmove 12 from 5 to 8\nmove 3 from 7 to 9\nmove 18 from 8 to 9\nmove 2 from 8 to 6\nmove 3 from 2 to 3\nmove 14 from 9 to 4\nmove 1 from 1 to 3\nmove 7 from 9 to 3\nmove 1 from 2 to 1\nmove 8 from 4 to 5\nmove 5 from 6 to 3\nmove 2 from 7 to 9\nmove 3 from 4 to 9\nmove 4 from 9 to 6\nmove 4 from 6 to 1\nmove 8 from 4 to 6\nmove 10 from 1 to 2\nmove 13 from 3 to 2\nmove 17 from 5 to 9\nmove 2 from 5 to 1\nmove 9 from 9 to 7\nmove 1 from 3 to 6\nmove 2 from 1 to 8\nmove 11 from 2 to 4\nmove 5 from 6 to 8\nmove 1 from 6 to 3\nmove 1 from 1 to 4\nmove 3 from 8 to 6\nmove 3 from 2 to 8\nmove 9 from 7 to 9\nmove 4 from 4 to 7\nmove 1 from 9 to 5\nmove 15 from 9 to 7\nmove 7 from 8 to 3\nmove 1 from 5 to 6\nmove 2 from 6 to 9\nmove 8 from 2 to 6\nmove 3 from 4 to 3\nmove 1 from 2 to 5\nmove 4 from 9 to 3\nmove 1 from 3 to 4\nmove 13 from 6 to 2\nmove 1 from 5 to 1\nmove 4 from 4 to 9\nmove 6 from 3 to 2\nmove 11 from 2 to 7\nmove 6 from 3 to 4\nmove 3 from 3 to 2\nmove 1 from 3 to 4\nmove 1 from 1 to 3\nmove 3 from 9 to 2\nmove 1 from 3 to 1\nmove 4 from 7 to 1\nmove 1 from 9 to 5\nmove 5 from 1 to 4\nmove 11 from 2 to 4\nmove 1 from 5 to 3\nmove 1 from 2 to 3\nmove 12 from 4 to 2\nmove 2 from 7 to 2\nmove 7 from 4 to 3\nmove 5 from 4 to 1\nmove 7 from 7 to 6\nmove 4 from 1 to 8\nmove 1 from 8 to 5\nmove 8 from 3 to 2\nmove 4 from 7 to 4\nmove 13 from 7 to 1\nmove 2 from 8 to 6\nmove 5 from 4 to 9\nmove 1 from 3 to 6\nmove 1 from 5 to 8\nmove 1 from 2 to 9\nmove 4 from 2 to 6\nmove 2 from 8 to 6\nmove 10 from 1 to 3\nmove 4 from 9 to 4\nmove 2 from 1 to 3\nmove 5 from 2 to 9\nmove 4 from 9 to 2\nmove 1 from 1 to 2\nmove 13 from 2 to 4\nmove 15 from 4 to 5\nmove 3 from 6 to 8\nmove 8 from 3 to 8\nmove 1 from 4 to 2\nmove 14 from 5 to 1\nmove 1 from 5 to 4\nmove 1 from 4 to 2\nmove 8 from 6 to 7\nmove 3 from 6 to 2\nmove 2 from 9 to 1\nmove 8 from 8 to 7\nmove 9 from 1 to 5\nmove 7 from 5 to 3\nmove 14 from 7 to 9\nmove 2 from 2 to 3\nmove 7 from 2 to 1\nmove 1 from 6 to 1\nmove 4 from 9 to 2\nmove 8 from 3 to 6\nmove 2 from 4 to 3\nmove 4 from 3 to 5\nmove 5 from 5 to 7\nmove 2 from 6 to 9\nmove 6 from 6 to 2\nmove 4 from 2 to 3\nmove 1 from 6 to 2\nmove 2 from 7 to 8\nmove 13 from 9 to 5\nmove 2 from 7 to 1\nmove 14 from 1 to 5\nmove 15 from 5 to 7\nmove 3 from 8 to 7\nmove 5 from 3 to 5\nmove 6 from 5 to 7\nmove 4 from 1 to 7\nmove 1 from 2 to 5\nmove 3 from 2 to 8\nmove 11 from 5 to 2\nmove 10 from 7 to 1\nmove 1 from 3 to 4\nmove 10 from 2 to 9\nmove 1 from 5 to 8\nmove 6 from 7 to 3\nmove 1 from 4 to 6\nmove 2 from 3 to 8\nmove 1 from 2 to 1\nmove 4 from 3 to 9\nmove 3 from 1 to 6\nmove 2 from 7 to 1\nmove 1 from 5 to 6\nmove 1 from 3 to 8\nmove 4 from 1 to 4\nmove 5 from 2 to 9\nmove 3 from 1 to 4\nmove 18 from 9 to 7\nmove 4 from 8 to 4\nmove 3 from 1 to 2\nmove 1 from 9 to 7\nmove 1 from 4 to 7\nmove 1 from 6 to 2\nmove 1 from 2 to 5\nmove 25 from 7 to 3\nmove 7 from 4 to 2\nmove 8 from 7 to 9\nmove 4 from 8 to 6\nmove 1 from 8 to 5\nmove 4 from 6 to 5\nmove 2 from 9 to 5\nmove 3 from 5 to 8\nmove 4 from 6 to 4\nmove 12 from 3 to 5\nmove 11 from 3 to 2\nmove 13 from 5 to 8\nmove 4 from 9 to 6\nmove 7 from 4 to 9\nmove 2 from 6 to 2\nmove 12 from 2 to 7\nmove 1 from 6 to 3\nmove 1 from 5 to 6\nmove 2 from 5 to 3\nmove 15 from 8 to 6\nmove 4 from 6 to 7\nmove 1 from 5 to 1\nmove 10 from 2 to 8\nmove 8 from 8 to 3\nmove 8 from 6 to 8\nmove 2 from 7 to 6\nmove 9 from 9 to 7\nmove 8 from 8 to 9\nmove 1 from 1 to 3\nmove 1 from 2 to 7\nmove 7 from 3 to 1\nmove 3 from 8 to 5\nmove 3 from 1 to 6\nmove 7 from 9 to 2\nmove 2 from 3 to 7\nmove 5 from 7 to 9\nmove 17 from 7 to 5\nmove 2 from 7 to 6\nmove 10 from 6 to 3\nmove 1 from 1 to 3\nmove 6 from 9 to 3\nmove 1 from 2 to 9\nmove 2 from 7 to 9\nmove 2 from 9 to 7\nmove 1 from 5 to 8\nmove 1 from 8 to 5\nmove 6 from 2 to 5\nmove 1 from 6 to 1\nmove 5 from 3 to 5\nmove 1 from 6 to 8\nmove 1 from 7 to 9\nmove 2 from 9 to 3\nmove 15 from 5 to 2\nmove 2 from 1 to 8\nmove 2 from 3 to 7\nmove 2 from 8 to 3\nmove 3 from 5 to 9\nmove 1 from 8 to 6\nmove 1 from 9 to 6\nmove 3 from 7 to 6\nmove 17 from 3 to 4\nmove 1 from 1 to 2\nmove 6 from 2 to 9\nmove 16 from 4 to 1\nmove 4 from 6 to 8\nmove 9 from 5 to 6\nmove 8 from 6 to 2\nmove 2 from 9 to 5\nmove 2 from 3 to 5\nmove 1 from 6 to 2\nmove 1 from 4 to 8\nmove 14 from 1 to 3\nmove 8 from 5 to 3\nmove 20 from 3 to 1\nmove 1 from 8 to 2\nmove 1 from 9 to 6\nmove 1 from 6 to 7\nmove 1 from 7 to 3\nmove 22 from 1 to 2\nmove 3 from 3 to 6\nmove 27 from 2 to 8\nmove 2 from 2 to 8\nmove 2 from 6 to 9\nmove 2 from 9 to 4\nmove 2 from 4 to 8\nmove 1 from 1 to 3\nmove 14 from 8 to 5\nmove 1 from 3 to 9\nmove 3 from 9 to 2\nmove 5 from 2 to 8\nmove 10 from 2 to 9\nmove 1 from 6 to 7\nmove 1 from 7 to 5\nmove 7 from 5 to 2\nmove 2 from 9 to 2\nmove 1 from 6 to 2\nmove 2 from 9 to 5\nmove 3 from 5 to 6\nmove 6 from 5 to 3\nmove 1 from 5 to 6\nmove 4 from 3 to 9\nmove 2 from 9 to 8\nmove 3 from 9 to 5\nmove 23 from 8 to 1\nmove 2 from 6 to 1\nmove 1 from 5 to 7\nmove 2 from 3 to 5\nmove 2 from 9 to 5\nmove 4 from 9 to 7\nmove 2 from 9 to 4\nmove 1 from 5 to 4\nmove 5 from 8 to 5\nmove 2 from 6 to 2\nmove 3 from 7 to 3\nmove 1 from 3 to 4\nmove 3 from 2 to 8\nmove 4 from 1 to 6\nmove 2 from 6 to 3\nmove 4 from 1 to 2\nmove 3 from 8 to 1\nmove 13 from 2 to 5\nmove 4 from 3 to 2\nmove 14 from 5 to 7\nmove 5 from 2 to 7\nmove 18 from 7 to 9\nmove 4 from 4 to 7\nmove 2 from 5 to 4\nmove 17 from 9 to 5\nmove 1 from 9 to 1\nmove 1 from 7 to 2\nmove 5 from 7 to 2\nmove 18 from 1 to 4\nmove 1 from 7 to 3\nmove 1 from 3 to 6\nmove 2 from 1 to 3\nmove 1 from 6 to 5\nmove 2 from 6 to 8\nmove 1 from 8 to 9\nmove 1 from 8 to 3\nmove 13 from 4 to 5\nmove 1 from 1 to 6\nmove 3 from 2 to 4\nmove 1 from 6 to 1\nmove 3 from 2 to 9\nmove 3 from 3 to 1\nmove 5 from 4 to 5\nmove 30 from 5 to 3\nmove 1 from 4 to 6\nmove 1 from 9 to 8\nmove 1 from 9 to 6\nmove 21 from 3 to 7\nmove 3 from 1 to 6\nmove 1 from 1 to 4\nmove 1 from 9 to 6\nmove 1 from 8 to 2\nmove 1 from 3 to 6\nmove 1 from 9 to 3\nmove 5 from 4 to 8\nmove 1 from 2 to 4\nmove 9 from 5 to 7\nmove 2 from 5 to 9\nmove 2 from 8 to 2\nmove 2 from 6 to 3\nmove 1 from 4 to 1\nmove 4 from 3 to 8\nmove 2 from 9 to 2\nmove 4 from 2 to 6\nmove 1 from 1 to 4\nmove 2 from 6 to 9\nmove 2 from 5 to 4\nmove 1 from 3 to 1\nmove 1 from 1 to 3\nmove 2 from 9 to 1\nmove 5 from 3 to 5\nmove 1 from 1 to 8\nmove 4 from 6 to 4\nmove 5 from 5 to 6\nmove 18 from 7 to 5\nmove 1 from 3 to 4\nmove 12 from 7 to 5\nmove 15 from 5 to 6\nmove 1 from 5 to 8\nmove 1 from 3 to 7\nmove 1 from 1 to 2\nmove 1 from 2 to 4\nmove 1 from 7 to 9\nmove 2 from 8 to 2\nmove 1 from 2 to 4\nmove 4 from 4 to 2\nmove 1 from 2 to 1\nmove 1 from 9 to 8\nmove 4 from 6 to 4\nmove 3 from 2 to 6\nmove 1 from 2 to 6\nmove 8 from 4 to 3\nmove 1 from 1 to 3\nmove 6 from 6 to 1\nmove 1 from 3 to 6\nmove 5 from 1 to 7\nmove 10 from 5 to 9\nmove 3 from 9 to 8\nmove 7 from 6 to 2\nmove 1 from 7 to 8\nmove 3 from 5 to 8\nmove 3 from 6 to 2\nmove 6 from 8 to 9\nmove 1 from 5 to 3\nmove 2 from 3 to 1\nmove 2 from 4 to 8\nmove 6 from 6 to 9\nmove 1 from 1 to 4\nmove 17 from 9 to 2\nmove 1 from 4 to 1\nmove 2 from 7 to 8\nmove 1 from 9 to 8\nmove 3 from 8 to 4\nmove 3 from 1 to 4\nmove 9 from 8 to 2\nmove 1 from 8 to 4\nmove 12 from 2 to 7\nmove 4 from 7 to 4\nmove 1 from 8 to 1\nmove 10 from 4 to 2\nmove 3 from 3 to 2\nmove 1 from 9 to 7\nmove 11 from 7 to 3\nmove 1 from 3 to 1\nmove 2 from 3 to 9\nmove 1 from 3 to 7\nmove 2 from 1 to 9\nmove 1 from 6 to 5\nmove 7 from 3 to 6\nmove 1 from 7 to 3\nmove 3 from 3 to 4\nmove 1 from 5 to 7\nmove 2 from 4 to 3\nmove 2 from 4 to 8\nmove 1 from 7 to 6\nmove 2 from 6 to 8\nmove 1 from 9 to 2\nmove 1 from 9 to 5\nmove 1 from 5 to 1\nmove 1 from 8 to 6\nmove 1 from 3 to 2\nmove 4 from 6 to 1\nmove 5 from 1 to 4\nmove 11 from 2 to 4\nmove 2 from 8 to 2\nmove 1 from 8 to 9\nmove 27 from 2 to 5\nmove 4 from 6 to 3\nmove 3 from 2 to 4\nmove 2 from 5 to 9\nmove 1 from 5 to 7\nmove 2 from 9 to 5\nmove 14 from 4 to 7\nmove 2 from 4 to 7\nmove 3 from 4 to 8\nmove 4 from 3 to 1\nmove 4 from 1 to 8\nmove 2 from 3 to 9\nmove 2 from 9 to 3\nmove 7 from 8 to 9\nmove 1 from 3 to 8\nmove 2 from 3 to 2\nmove 25 from 5 to 9\nmove 1 from 5 to 8\nmove 1 from 8 to 7\nmove 26 from 9 to 1\nmove 23 from 1 to 5\nmove 7 from 9 to 7\nmove 1 from 9 to 8\nmove 1 from 9 to 2\nmove 5 from 7 to 1\nmove 20 from 5 to 6\nmove 1 from 7 to 6\nmove 2 from 5 to 3\nmove 1 from 8 to 6\nmove 21 from 6 to 8\nmove 1 from 6 to 4\nmove 1 from 1 to 7\nmove 2 from 1 to 6\nmove 1 from 1 to 3\nmove 1 from 2 to 5\nmove 1 from 2 to 6\nmove 2 from 7 to 6\nmove 6 from 7 to 9\nmove 3 from 1 to 2\nmove 17 from 8 to 1\nmove 1 from 4 to 1\nmove 2 from 6 to 9\nmove 3 from 8 to 9\nmove 2 from 3 to 7\nmove 2 from 9 to 8\nmove 4 from 7 to 3\nmove 4 from 3 to 4\nmove 2 from 5 to 8\nmove 4 from 8 to 4\nmove 3 from 6 to 8\nmove 18 from 1 to 5\nmove 1 from 3 to 4\nmove 3 from 2 to 4\nmove 5 from 9 to 1\nmove 10 from 7 to 5\nmove 5 from 1 to 3\nmove 5 from 3 to 5\nmove 5 from 4 to 3\nmove 2 from 4 to 2\nmove 5 from 8 to 3\nmove 25 from 5 to 2\nmove 3 from 3 to 6\nmove 1 from 1 to 3\nmove 3 from 6 to 7\nmove 1 from 4 to 2\nmove 1 from 5 to 8\nmove 2 from 4 to 9\nmove 1 from 8 to 1\nmove 20 from 2 to 7\nmove 10 from 7 to 1\nmove 1 from 1 to 7\nmove 4 from 7 to 8\nmove 5 from 5 to 4\nmove 4 from 8 to 6\nmove 1 from 1 to 3\nmove 5 from 7 to 4\nmove 2 from 1 to 5\nmove 4 from 9 to 1\nmove 3 from 2 to 5\nmove 5 from 5 to 1\nmove 1 from 9 to 1\nmove 11 from 1 to 3\nmove 1 from 6 to 2\nmove 7 from 3 to 5\nmove 11 from 3 to 7\nmove 1 from 2 to 6\nmove 7 from 7 to 8\nmove 1 from 9 to 1\nmove 2 from 3 to 1\nmove 1 from 5 to 3\nmove 4 from 1 to 6\nmove 4 from 6 to 3\nmove 9 from 4 to 5\nmove 2 from 8 to 2\nmove 4 from 6 to 9\nmove 3 from 2 to 4\nmove 1 from 8 to 6"

func reversed(s []string) []string {
	// yoink'ed this one from SO
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Part1(rev ...bool) string {
	reverse := true // true for part1, false for part2

	if len(rev) > 0 {
		reverse = rev[0]
	}
	m := make(map[int][]string)

	for index, line := range strings.Split(input, "\n") {
		if index <= 7 {
			var seq int
			for i := 1; i < len(line); i += 4 {
				seq++
				q := string(line[i])
				if q == " " {
					continue
				}
				if val, ok := m[seq]; ok {
					m[seq] = append(val, q)
				} else {
					m[seq] = []string{q}
				}
			}
		}
		if index > 9 {
			var from, to, count int
			_, _ = fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
			var t []string
			if reverse {
				t = append(t, reversed(m[from][:count])...)
			} else {
				t = append(t, m[from][:count]...)
			}
			t = append(t, m[to]...)
			m[to] = t
			m[from] = append(m[from][:0], m[from][count:]...)
		}
	}

	var q strings.Builder
	for i := 1; i < 10; i++ {
		q.Write([]byte(m[i][0]))
	}
	return q.String()
}

func Part2() string {
	return Part1(false)
}
