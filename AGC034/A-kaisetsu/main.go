package main

import (
	"fmt"
)

func main() {
	var n int
	var a int
	var b int
	var c int
	var d int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Scan(&s)

	ret := "No"
	if c < d {
		sharpTwoFlg := false
		sharpCount := 0
		for i := a; i < c; i++ {
			if s[i] == '#' {
				sharpCount++
				if sharpCount >= 2 {
					sharpTwoFlg = true
					break
				}
			} else {
				sharpCount = 0
			}
		}

		sharpCount = 0
		for i := b; i < d; i++ {
			if s[i] == '#' {
				sharpCount++
				if sharpCount >= 2 {
					sharpTwoFlg = true
					break
				}
			} else {
				sharpCount = 0
			}
		}

		if sharpTwoFlg {
			ret = "No"
		} else {
			ret = "Yes"
		}

	} else {

		dotThreeFlg := false
		sharpTwoFlg := false

		dotCount := 0
		sharpCount := 0
		for i := b - 2; i <= d; i++ {
			if s[i] == '.' {
				dotCount++
				if dotCount >= 3 {
					dotThreeFlg = true
				}
				sharpCount = 0
			} else {
				sharpCount++
				if sharpCount >= 2 {
					sharpTwoFlg = true
					break
				}
				dotCount = 0
			}
		}

		sharpCount = 0
		for i := a; i < c; i++ {
			if s[i] == '#' {
				sharpCount++
				if sharpCount >= 2 {
					sharpTwoFlg = true
					break
				}
			} else {
				sharpCount = 0
			}
		}

		if !sharpTwoFlg && dotThreeFlg {
			ret = "Yes"
		}
	}

	fmt.Println(ret)
}
