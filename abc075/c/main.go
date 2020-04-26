package main

import (
	"bufio"
	"fmt"
	"os"
)

var graph [60][60]bool
var visited []bool

type edge struct {
	a int
	b int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)

	var a, b int
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		a--
		b--

		graph[a][b] = true
		graph[b][a] = true
		edges[i] = edge{a: a, b: b}
	}

	ans := 0
	for _, e := range edges {
		// 辺を切ってみる
		graph[e.a][e.b] = false
		graph[e.b][e.a] = false

		// DFSで連結されているかをチェック
		visited = make([]bool, n)
		dfs(0, n)
		for i := 0; i < n; i++ {
			if !visited[i] {
				ans++
				break
			}
		}

		// 切った辺を戻しておく
		graph[e.a][e.b] = true
		graph[e.b][e.a] = true
	}

	fmt.Println(ans)
}

func dfs(x, n int) {
	if visited[x] {
		// すでの訪れている場合は探索終了
		return
	}

	// 連結されていることを記録
	visited[x] = true

	// 辺でつながっている次のノードを探索
	for i := 0; i < n; i++ {
		if graph[x][i] {
			dfs(i, n)
		}
	}
}

// Utility

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

func primeFactor(x int) map[int]int {
	res := make(map[int]int)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			res[i]++
			x = x / i
		}
	}
	if x != 1 {
		res[x] = 1
	}
	return res
}

func divisor(x int) []int {
	res := make([]int, 0)
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			res = append(res, i)
			if i != x/i {
				res = append(res, x/i)
			}
		}
	}
	return res
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm(x, y int) int {
	return x / gcd(x, y) * y
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
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

// Stack ...
type Stack []int

// pop ...
func (s *Stack) empty() bool {
	return len(*s) == 0
}

// push ...
func (s *Stack) push(i int) {
	*s = append(*s, i)
}

// pop ...
func (s *Stack) pop() (int, bool) {
	if s.empty() {
		return 0, false
	} else {
		index := len(*s) - 1
		res := (*s)[index]
		*s = (*s)[:index]
		return res, true
	}
}
