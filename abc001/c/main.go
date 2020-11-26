package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var MOD = 1000000007
var INF = 1001001001

func main() {
	r := bufio.NewReader(os.Stdin)
	var deg, dis float64
	fmt.Fscan(r, &deg)
	fmt.Fscan(r, &dis)

	dir := "N"
	w := 0
	if 113 <= deg && deg < 338 {
		dir = "NNE"
	} else if 338 <= deg && deg < 563 {
		dir = "NE"
	} else if 563 <= deg && deg < 788 {
		dir = "ENE"
	} else if 788 <= deg && deg < 1013 {
		dir = "E"
	} else if 1013 <= deg && deg < 1238 {
		dir = "ESE"
	} else if 1238 <= deg && deg < 1463 {
		dir = "SE"
	} else if 1463 <= deg && deg < 1688 {
		dir = "SSE"
	} else if 1688 <= deg && deg < 1913 {
		dir = "S"
	} else if 1913 <= deg && deg < 2138 {
		dir = "SSW"
	} else if 2138 <= deg && deg < 2363 {
		dir = "SW"
	} else if 2363 <= deg && deg < 2588 {
		dir = "WSW"
	} else if 2588 <= deg && deg < 2813 {
		dir = "W"
	} else if 2813 <= deg && deg < 3038 {
		dir = "WNW"
	} else if 3038 <= deg && deg < 3263 {
		dir = "NW"
	} else if 3263 <= deg && deg < 3488 {
		dir = "NNW"
	} else {
		dir = "N"
	}

	disfloat := math.Round(dis/60*10) / 10
	if 0.0 <= disfloat && disfloat <= 0.2 {
		w = 0
	} else if 0.3 <= disfloat && disfloat <= 1.5 {
		w = 1
	} else if 1.6 <= disfloat && disfloat <= 3.3 {
		w = 2
	} else if 3.4 <= disfloat && disfloat <= 5.4 {
		w = 3
	} else if 5.5 <= disfloat && disfloat <= 7.9 {
		w = 4
	} else if 8.0 <= disfloat && disfloat <= 10.7 {
		w = 5
	} else if 10.8 <= disfloat && disfloat <= 13.8 {
		w = 6
	} else if 13.9 <= disfloat && disfloat <= 17.1 {
		w = 7
	} else if 17.2 <= disfloat && disfloat <= 20.7 {
		w = 8
	} else if 20.8 <= disfloat && disfloat <= 24.4 {
		w = 9
	} else if 24.5 <= disfloat && disfloat <= 28.4 {
		w = 10
	} else if 28.5 <= disfloat && disfloat <= 32.6 {
		w = 11
	} else {
		w = 12
	}

	if w == 0 {
		fmt.Println("C", 0)
	} else {
		fmt.Println(dir, w)
	}
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
func (q *Queue) pop() int {
	if q.empty() {
		return 0
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res
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
func (s *Stack) pop() int {
	if s.empty() {
		return 0
	}
	index := len(*s) - 1
	res := (*s)[index]
	*s = (*s)[:index]
	return res

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
