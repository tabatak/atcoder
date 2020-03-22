package main

import (
	"bufio"
	"fmt"
	"math"
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

	dy := []int{1, 0}
	dx := []int{0, 1}

	cost := make([][]int, h)
	for i := 0; i < h; i++ {
		cost[i] = make([]int, w)
		for j := 0; j < w; j++ {
			cost[i][j] = math.MaxInt64
		}
	}
	if grid[0][0] == '#' {
		cost[0][0] = 1
	} else {
		cost[0][0] = 0
	}

	var q Queue
	// スタート地点をpush
	q.push(Point{0, 0})

	for !q.empty() {
		p, _ := q.pop()
		x := p.x
		y := p.y

		for i := 0; i < 2; i++ {
			ny := y + dy[i]
			nx := x + dx[i]

			if ny >= h || nx >= w {
				//グリッド外は探索不要
				continue
			}

			if cost[y][x] >= cost[ny][nx] {
				//コストが最小にならない経路は探索不要
				continue
			}

			tmpCost := cost[y][x]
			if grid[y][x] == '.' && grid[ny][nx] == '#' {
				// 白->黒への移動の場合はコスト+1
				tmpCost++
			}
			cost[ny][nx] = tmpCost

			q.push(Point{nx, ny})
		}

	}

	fmt.Println(cost[h-1][w-1])
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
