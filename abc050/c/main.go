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
	as := make([]int, n)
	pos := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])
		pos[as[i]]++
	}

	// 証言が正しいかをチェック
	flg := true
	for i := n - 1; i >= 0; i -= 2 {
		count := pos[i]
		if i == 0 {
			if count != 1 {
				flg = false
				break
			}
		} else {
			if count != 2 {
				flg = false
				break
			}
		}
	}
	if !flg {
		fmt.Println(0)
		return
	}

	mod := 1000000007
	ans := modpow(2, n/2, mod)
	fmt.Println(ans)
}

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
