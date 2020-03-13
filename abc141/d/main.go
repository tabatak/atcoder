package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

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

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)

	// heap
	ih := &intHeap{}
	heap.Init(ih)

	var a int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a)
		heap.Push(ih, a)
	}

	// 割引券を値段の高い順に適用する
	for m > 0 {
		// 最も高い値段を取り出す
		price := heap.Pop(ih).(int)
		// 半額クーポンを適用
		heap.Push(ih, price/2)
		m--
	}

	ans := 0
	for ih.Len() > 0 {
		ans += heap.Pop(ih).(int)
	}

	fmt.Println(ans)
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
