package main

import (
	"bufio"
	"fmt"
	"os"
)

type testimony struct {
	x int
	y int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	var a, x, y int
	ts := make([][]testimony, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a)
		testimonies := make([]testimony, a)
		for j := 0; j < a; j++ {
			fmt.Fscan(r, &x)
			fmt.Fscan(r, &y)
			testimonies[j] = testimony{x, y}
		}
		ts[i] = testimonies
	}

	ans := 0
	// bitで網羅してみる
	for i := 0; i < (1 << uint(n)); i++ {
		flg := true
		for j := 0; j < n; j++ {
			// 証言者が正直者かどうか
			honest := i>>uint(j)&1 == 1
			if !honest {
				// 不親切な人ならば証言のチェックはしない
				continue
			}
			for _, t := range ts[j] {
				if i>>uint(t.x-1)&1 != t.y {
					// 正直者の言う通りでなければまちがい
					flg = false
					break
				}
			}

			if !flg {
				break
			}
		}
		if flg {
			count := 0
			for j := 0; j < n; j++ {
				count += i >> uint(j) & 1
			}

			ans = max(ans, count)
		}
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
