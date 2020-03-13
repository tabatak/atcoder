package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s[i])
	}
}

// Point ...
type Point struct {
	x int
	y int
}

// Queue ...
type Queue []Point

// pop ...
func (q *Queue) empty() bool {
	return len(*q) == 0
}

// push ...
func (q *Queue) push(i Point) {
	*q = append(*q, i)
}

// pop ...
func (q *Queue) pop() (Point, bool) {
	if q.empty() {
		return Point{0, 0}, false
	} else {
		res := (*q)[0]
		*q = (*q)[1:]
		return res, true
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
