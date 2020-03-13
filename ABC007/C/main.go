package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var row, col, sx, sy, gx, gy int
	fmt.Fscan(r, &row)
	fmt.Fscan(r, &col)
	fmt.Fscan(r, &sy)
	fmt.Fscan(r, &sx)
	fmt.Fscan(r, &gy)
	fmt.Fscan(r, &gx)
	sx--
	sy--
	gx--
	gy--

	s := make([]string, row)
	for i := 0; i < row; i++ {
		fmt.Fscan(r, &s[i])
	}

	// 上下左右への移動
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	// 各ポイントへのコスト
	d := make([][]int, row)
	for i := 0; i < row; i++ {
		c := make([]int, col)
		for j := 0; j < col; j++ {
			c[j] = math.MaxInt64
		}
		d[i] = c
	}
	// スタートポイントのコストは0
	d[sy][sx] = 0

	var q Queue
	q.push(Point{sx, sy})

	for !q.empty() {
		p, _ := q.pop()
		if p.x == gx && p.y == gy {
			break
		}

		for i := 0; i < 4; i++ {
			//上下左右の探索対象を探す
			nx := p.x + dx[i]
			ny := p.y + dy[i]

			if 0 <= nx && nx < col && 0 <= ny && ny < row && s[ny][nx] != '#' && d[ny][nx] == math.MaxInt64 {
				// 探索対象ならキューに追加
				q.push(Point{nx, ny})
				// ポイントへのコストは現在ポイントまでのコスト+1
				d[ny][nx] = d[p.y][p.x] + 1
			}
		}
	}

	// ゴールポイントのコストを出力
	fmt.Println(d[gy][gx])
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
