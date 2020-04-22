package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to   int
	cost int
}

var tree [100010][]edge
var depth [100010]int

func dfs(v, p, d int) {
	depth[v] = d
	for _, e := range tree[v] {
		if e.to == p {
			continue
		}
		dfs(e.to, v, d+e.cost)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)

	// 隣接リスト
	var a, b, c int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		fmt.Fscan(r, &c)
		a--
		b--
		tree[a] = append(tree[a], edge{to: b, cost: c})
		tree[b] = append(tree[b], edge{to: a, cost: c})
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var q, k int
	fmt.Fscan(r, &q)
	fmt.Fscan(r, &k)
	k--
	dfs(k, -1, 0)

	var x, y int
	for i := 0; i < q; i++ {
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		x--
		y--

		fmt.Fprintln(w, depth[x]+depth[y])
	}
}

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	var n int
// 	fmt.Fscan(r, &n)

// 	// グラフをつくってみる
// 	var a, b, c int
// 	g := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		g[i] = make([]int, n+1)
// 	}
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(r, &a)
// 		fmt.Fscan(r, &b)
// 		fmt.Fscan(r, &c)
// 		g[a][b] = c
// 		g[b][a] = c
// 	}

// 	w := bufio.NewWriter(os.Stdout)
// 	defer w.Flush()

// 	var q, k int
// 	fmt.Fscan(r, &q)
// 	fmt.Fscan(r, &k)
// 	var x, y int
// 	for i := 0; i < q; i++ {
// 		fmt.Fscan(r, &x)
// 		fmt.Fscan(r, &y)

// 		// x -> kの最短経路
// 		var q Queue
// 		q.push(x)
// 		for !q.empty() {
// 			current, _ := q.pop()
// 			if current != k {
// 				for j := 1; j < n; j++ {
// 					if g[x][j] == 0 {
// 						continue
// 					}
// 					q.push(j)
// 					g[x][j] = g[x][current] + g[current][j]
// 					g[j][x] = g[x][current] + g[current][j]
// 				}
// 			}
// 		}

// 		// k -> yの最短経路
// 		q.push(k)
// 		for !q.empty() {
// 			current, _ := q.pop()
// 			if current != y {
// 				for j := 1; j < n; j++ {
// 					if g[k][j] == 0 {
// 						continue
// 					}
// 					q.push(j)
// 					g[k][j] = g[k][current] + g[current][j]
// 					g[j][k] = g[k][current] + g[current][j]
// 				}
// 			}
// 		}

// 		fmt.Fprintln(w, g[x][k]+g[k][y])
// 	}
// }

// Union-Find
type unionFind struct {
	d []int
}

func newUnionFind(n int) *unionFind {
	uf := new(unionFind)
	uf.d = make([]int, n)
	for i := 0; i < n; i++ {
		uf.d[i] = -1
	}
	return uf
}

func (uf *unionFind) find(x int) int {
	if uf.d[x] < 0 {
		return x
	}
	uf.d[x] = uf.find(uf.d[x])
	return uf.d[x]
}

func (uf *unionFind) unite(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return false
	}

	if uf.d[rootX] < uf.d[rootY] {
		uf.d[rootX] += uf.d[rootY]
		uf.d[rootY] = rootX
	} else {
		uf.d[rootY] += uf.d[rootX]
		uf.d[rootX] = rootY
	}

	return true
}

func (uf *unionFind) same(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *unionFind) size(x int) int {
	return -uf.d[uf.find(x)]
}

// mod combination
func modpow(a, n, mod int) int {
	res := 1
	for n > 0 {
		if n%2 != 0 {
			res = res * a % mod
		}
		a = a * a % mod
		n = n / 2
	}
	return res
}

func modcomb(n, a, mod int) int {
	x := 1
	y := 1
	for i := 0; i < a; i++ {
		x = x * (n - i)
		x %= mod
		y = y * (i + 1)
		y %= mod
	}
	return x * modpow(y, mod-2, mod) % mod
}

func modfactorial(n, mod int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = (result * i) % mod
	}
	return result
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
