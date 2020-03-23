package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var h, w, k int
	fmt.Fscan(r, &h)
	fmt.Fscan(r, &w)
	fmt.Fscan(r, &k)

	grid := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &grid[i])
	}

	ans := math.MaxInt64

	// 水平方向の切り方を網羅する（h <= 10）
	for div := 0; div < (1 << uint(h-1)); div++ {
		// 水平方向に分割されるグループ数
		g := 0
		id := make([]int, h)

		for i := 0; i < h; i++ {
			id[i] = g
			if div>>uint(i)&1 == 1 {
				// 水平方向に分割する場合はグループを加算
				g++
			}
		}
		g++

		// グループ化されたグリッドの状態を一次元に変換
		c := make([][]int, g)
		for i := 0; i < g; i++ {
			c[i] = make([]int, w)
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				c[id[i]][j] += int(grid[i][j] - '0')
			}
		}

		num := g - 1
		now := make([]int, g)
		add := func(j int) bool {
			// WCの個数がK個を超えるグループがある場合はfalse
			for i := 0; i < g; i++ {
				now[i] += c[i][j]
			}
			flg := true
			for i := 0; i < g; i++ {
				if now[i] > k {
					flg = false
					break
				}
			}
			return flg
		}

		for j := 0; j < w; j++ {
			if !add(j) {
				num++
				now = make([]int, g)

				if !add(j) {
					// これ以上垂直方向での分割はできない
					num = math.MaxInt64
					break
				}
			}
		}

		ans = min(ans, num)
	}

	fmt.Println(ans)
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
