package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)

	ans := 0
	ans += n * (n - 1) / 2
	ans += m * (m - 1) / 2

	fmt.Println(ans)
}
