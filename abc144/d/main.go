package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a, b, x float64
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &x)

	s := x / a

	if s >= a*b/2 {
		h := (a*b - s) * 2 / a
		r := math.Atan2(h, a)
		fmt.Printf("%.9f\n", r/(2*math.Pi)*360)
	} else {
		w := s * 2 / b
		r := math.Atan2(b, w)
		fmt.Printf("%.9f\n", r/(2*math.Pi)*360)
	}
}
