package main

import "fmt"

// LeastInterval returns the fewest CPU slots to run all tasks when equal
// tasks must sit at least n slots apart (idle slots allowed). Closed form:
// the most frequent task builds the frame, everything else fills it.
func LeastInterval(tasks []byte, n int) int {
	var count [26]int
	for _, t := range tasks {
		count[t-'A']++
	}
	maxFreq, maxCount := 0, 0
	for _, c := range count {
		switch {
		case c > maxFreq:
			maxFreq, maxCount = c, 1
		case c == maxFreq && c > 0:
			maxCount++
		}
	}
	// (maxFreq-1) full frames of width n+1, plus the final row of max-freq tasks.
	framed := (maxFreq-1)*(n+1) + maxCount
	if framed < len(tasks) {
		return len(tasks) // enough variety: no idle slots at all
	}
	return framed
}

func main() {
	fmt.Println(LeastInterval([]byte("AAABBB"), 2))
	fmt.Println(LeastInterval([]byte("AAABBB"), 0))
}
