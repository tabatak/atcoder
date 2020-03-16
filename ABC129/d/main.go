package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(r, &h)
	fmt.Fscan(r, &w)

	grid := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &grid[i])
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	ans := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '#' {
				continue
			}

			tmp := 0
			for k := 0; k < 4; k++ {
				nx := j + dx[k]
				ny := i + dy[k]
				for 0 <= nx && nx < w && 0 <= ny && ny < h && grid[ny][nx] != '#' {
					tmp++
					nx += dx[k]
					ny += dy[k]
				}
			}
			ans = max(ans, tmp)
		}
	}
	fmt.Println(ans + 1)
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
