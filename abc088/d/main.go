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
	whiteCount := 0
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &grid[i])
		for j := 0; j < w; j++ {
			if grid[i][j] == '.' {
				whiteCount++
			}
		}
	}

	// 経路のコスト
	cost := make([][]int, h)
	for i := 0; i < h; i++ {
		cost[i] = make([]int, w)
		for j := 0; j < w; j++ {
			cost[i][j] = -1
		}
	}
	cost[0][0] = 0

	// 最短経路をBFSで探索
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	var q Queue
	q.push(Point{0, 0})
	for !q.empty() {
		p, _ := q.pop()

		for i := 0; i < 4; i++ {
			nx := p.x + dx[i]
			ny := p.y + dy[i]

			if 0 <= nx && nx < w && 0 <= ny && ny < h && grid[ny][nx] != '#' && cost[ny][nx] == -1 {
				// 未踏かつ白マスの場合はコストを計算し、キューに追加
				cost[ny][nx] = cost[p.y][p.x] + 1
				q.push(Point{nx, ny})
			}
		}
	}

	if cost[h-1][w-1] == -1 {
		fmt.Println(-1)
	} else {
		ans := whiteCount - cost[h-1][w-1] - 1
		fmt.Println(ans)
	}

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

// Utility
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
