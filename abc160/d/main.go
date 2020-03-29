package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, x, y int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &x)
	fmt.Fscan(r, &y)
	x--
	y--

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		// 各ノードから他のノードへの最短距離を求める

		cost := make([]int, n)
		for j := 0; j < n; j++ {
			cost[j] = -1
		}
		cost[i] = 0
		var q Queue
		q.push(i)

		for !q.empty() {
			current, _ := q.pop()

			// 1つ前のノード
			if current-1 >= 0 && cost[current-1] == -1 {
				q.push(current - 1)
				cost[current-1] = cost[current] + 1
			}
			// 1つ後のノード
			if current+1 < n && cost[current+1] == -1 {
				q.push(current + 1)
				cost[current+1] = cost[current] + 1
			}
			// ショートカット
			if current == x && cost[y] == -1 {
				q.push(y)
				cost[y] = cost[current] + 1
			}
			if current == y && cost[x] == -1 {
				q.push(x)
				cost[x] = cost[current] + 1
			}
		}

		for j := 0; j < n; j++ {
			ans[cost[j]]++
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 1; i < n; i++ {
		fmt.Fprintln(w, ans[i]/2)
	}
}

// Queue ...
type Queue []int

// pop ...
func (q *Queue) empty() bool {
	return len(*q) == 0
}

// push ...
func (q *Queue) push(i int) {
	*q = append(*q, i)
}

// pop ...
func (q *Queue) pop() (int, bool) {
	if q.empty() {
		return 0, false
	} else {
		res := (*q)[0]
		*q = (*q)[1:]
		return res, true
	}
}
