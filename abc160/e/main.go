package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var x, y, a, b, c int
	fmt.Fscan(r, &x)
	fmt.Fscan(r, &y)
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)
	as := make([]int, a)
	bs := make([]int, b)
	cs := make([]int, c)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &as[i])
	}
	for i := 0; i < b; i++ {
		fmt.Fscan(r, &bs[i])
	}
	for i := 0; i < c; i++ {
		fmt.Fscan(r, &cs[i])
	}
	sort.Sort(sort.IntSlice(as))
	sort.Sort(sort.IntSlice(bs))
	sort.Sort(sort.IntSlice(cs))

	ans := 0
	sumX := 0
	sumY := 0
	asIndex := a - 1
	bsIndex := b - 1
	csIndex := c - 1
	for sumX+sumY < x+y {
		if sumX < x {
			if csIndex < 0 || as[asIndex] >= cs[csIndex] {
				ans += as[asIndex]
				asIndex--
				sumX++
			} else if bsIndex < 0 || as[asIndex] <= bs[bsIndex] {
				ans += cs[csIndex]
				csIndex--
				sumX++
			}
		}
		if sumY < y {
			if csIndex < 0 || bs[bsIndex] >= cs[csIndex] {
				ans += bs[bsIndex]
				bsIndex--
				sumY++
			} else if asIndex < 0 || bs[bsIndex] <= as[asIndex] {
				ans += cs[csIndex]
				csIndex--
				sumY++
			}
		}
	}

	fmt.Println(ans)
}
