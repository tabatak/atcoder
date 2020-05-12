package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to int
	id int
}

var g [100005][]edge
var ans [100005]int

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)

	var a, b int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		a--
		b--
		g[a] = append(g[a], edge{b, i})
		g[b] = append(g[b], edge{a, i})
	}

	// 塗分け
	dfs(0, -1, -1)

	colorNum := 0
	for i := 0; i < len(g); i++ {
		colorNum = max(colorNum, len(g[i]))
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, colorNum)
	for i := 0; i < n-1; i++ {
		fmt.Fprintln(w, ans[i])
	}
}

func dfs(curID, parentID, prevEdgeColor int) {
	color := 1
	for _, e := range g[curID] {
		if e.to == parentID {
			continue
		}
		if color == prevEdgeColor {
			color++
		}
		ans[e.id] = color
		dfs(e.to, curID, color)
		color++
	}
}

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	var n int
// 	fmt.Fscan(r, &n)

// 	nodes := make([]map[int]bool, n)
// 	for i := 0; i < n; i++ {
// 		nodes[i] = make(map[int]bool)
// 	}

// 	ans := make([]int, n-1)
// 	var colors []bool
// 	var a, b int
// 	for i := 0; i < n-1; i++ {
// 		fmt.Fscan(r, &a)
// 		fmt.Fscan(r, &b)
// 		a--
// 		b--

// 		color := -1
// 		for j := range colors {
// 			flg := true
// 			if _, ok := nodes[a][j+1]; ok {
// 				flg = false
// 			} else if _, ok := nodes[b][j+1]; ok {
// 				flg = false
// 			}
// 			if flg {
// 				color = j + 1
// 				break
// 			}
// 		}

// 		if color == -1 {
// 			colors = append(colors, true)
// 			color = len(colors)
// 		}
// 		ans[i] = color
// 		nodes[a][color] = true
// 		nodes[b][color] = true
// 	}

// 	w := bufio.NewWriter(os.Stdout)
// 	defer w.Flush()
// 	fmt.Fprintln(w, len(colors))
// 	for _, v := range ans {
// 		fmt.Fprintln(w, v)
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

// mod
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

func modinv(n, mod int) int {
	return modpow(n, mod-2, mod)
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

// Combination ...
type Combination struct {
	facts, invs []int
	mod         int
}

// NewCombination ...
func NewCombination(n, mod int) Combination {
	return Combination{
		facts: make([]int, n+1),
		invs:  make([]int, n+1),
		mod:   mod,
	}
}

func (cmb *Combination) calc(n, k int) int {
	ret := cmb.facts[n] * cmb.invs[k]
	ret %= cmb.mod
	ret = ret * cmb.invs[n-k]
	ret %= cmb.mod
	return ret
}

func (cmb *Combination) init(n int) {
	cmb.facts[0] = 1
	// 階乗を算出
	for i := 1; i <= n; i++ {
		cmb.facts[i] = cmb.facts[i-1] * i % cmb.mod
	}
	// 逆元を算出
	cmb.invs[n] = modinv(cmb.facts[n], cmb.mod)
	for i := n - 1; i >= 0; i-- {
		cmb.invs[i] = cmb.invs[i+1] * (i + 1) % cmb.mod
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
