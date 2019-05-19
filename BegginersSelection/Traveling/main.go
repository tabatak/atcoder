package main

import (
	"fmt"
	"math"
)

type loc struct {
	t float64
	x float64
	y float64
}

func main() {
	var n int
	fmt.Scan(&n)

	var a float64
	var b float64
	var c float64

	locs := make([]loc, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)

		locs[i] = loc{
			t: a,
			x: b,
			y: c,
		}
	}

	result := "Yes"
	for j := 0; j < n; j++ {
		traveling := float64(0)
		count := float64(0)
		if j == 0 {
			traveling = math.Abs(locs[j].x) + math.Abs(locs[j].y)
			count = locs[j].t
		} else {
			traveling = math.Abs(locs[j].x-(locs[j-1].x)) + math.Abs(locs[j].y-(locs[j-1].y))
			count = locs[j].t - locs[j-1].t
		}

		if traveling == count || (count > traveling && int(count-traveling)%2 == 0) {
			result = "Yes"
		} else {
			result = "No"
			break
		}
	}

	fmt.Println(result)
}
