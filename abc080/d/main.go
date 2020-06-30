package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var mod = 1000000007
var inf = 1000000007

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, c int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &c)

	prgs := points(make([]point, n))
	for i := 0; i < n; i++ {
		prg := point{}
		fmt.Fscan(r, &prg.s)
		fmt.Fscan(r, &prg.t)
		fmt.Fscan(r, &prg.c)
		prgs[i] = prg
	}
	// チャンネル、開始時間の昇順にソート
	sort.Sort(prgs)

	var merged points
	bc := prgs[0].c
	bs := prgs[0].s
	bt := prgs[0].t
	index := 1
	for index < n {
		if bc == prgs[index].c && bt == prgs[index].s {
			bt = prgs[index].t
		} else {
			merged = append(merged, point{bs, bt, bc})
			bs = prgs[index].s
			bt = prgs[index].t
			bc = prgs[index].c
		}

		if index == n-1 {
			merged = append(merged, point{bs, bt, bc})
		}

		index++
	}

	imos := make([]int, 100100)
	for _, v := range merged {
		imos[v.s-1]++
		imos[v.t]--
	}

	ans := 0
	tmp := 0
	for _, v := range imos {
		tmp += v
		ans = max(ans, tmp)
	}
	fmt.Println(ans)
}

// ↓嘘解法
// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	var n, c int
// 	fmt.Fscan(r, &n)
// 	fmt.Fscan(r, &c)

// 	prgs := points(make([]point, n))
// 	for i := 0; i < n; i++ {
// 		prg := point{}
// 		fmt.Fscan(r, &prg.s)
// 		fmt.Fscan(r, &prg.t)
// 		fmt.Fscan(r, &prg.c)
// 		prgs[i] = prg
// 	}
// 	// 開始時間の昇順にソート
// 	sort.Sort(prgs)

// 	// 番組を録画できるかをチェック
// 	recorders := make([]int, c)
// 	for _, prg := range prgs {
// 		for i := 0; i < c; i++ {
// 			if recorders[i] < prg.s {
// 				recorders[i] = prg.t
// 				break
// 			}
// 		}
// 	}

// 	// 使用した録画機の数を数える
// 	ans := 0
// 	for i := 0; i < c; i++ {
// 		if recorders[i] != 0 {
// 			ans++
// 		}
// 	}

// 	fmt.Println(ans)

// }

// sortable points
type point struct {
	s int
	t int
	c int
}

type points []point

func (p points) Len() int {
	return len(p)
}

func (p points) Less(i, j int) bool {
	if p[i].c != p[j].c {
		return p[i].c < p[j].c
	}
	return p[i].s < p[j].s
}

func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// permutations
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

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

func pow(a, n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= a
	}
	return ret
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

// priority_queue
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
