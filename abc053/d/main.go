package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	var a int
	nums := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a)
		nums[a]++
	}

	k := len(nums)

	if (n-k)%2 == 0 {
		fmt.Println(k)
	} else {
		fmt.Println(k - 1)
	}
}

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	var n int
// 	fmt.Fscan(r, &n)
// 	var a int
// 	nums := make(map[int]int)
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(r, &a)
// 		nums[a]++
// 	}

// 	// priority_queue
// 	ih := &intHeap{}
// 	heap.Init(ih)
// 	for _, v := range nums {
// 		if v >= 2 {
// 			heap.Push(ih, v)
// 		}
// 	}

// 	ans := n
// 	for ih.Len() > 0 {
// 		v1 := heap.Pop(ih)

// 		var v2 interface{}
// 		if ih.Len() > 0 {
// 			v2 = heap.Pop(ih)
// 		}

// 		if v1.(int) > 2 {
// 			heap.Push(ih, v1.(int)-1)
// 		}
// 		if v2 != nil && v2.(int) > 2 {
// 			heap.Push(ih, v2.(int)-1)
// 		}

// 		ans -= 2
// 	}

// 	if ans <= 1 {
// 		fmt.Println(1)
// 	} else {
// 		fmt.Println(ans)
// 	}
// }

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
