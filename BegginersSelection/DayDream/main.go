package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	sr := make([]byte, len(s))
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		sr[j] = s[i]
		j++
	}

	b := hasPrefix(string(sr))
	if b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func hasPrefix(s string) bool {
	if len(s) == 0 {
		return true
	}

	if strings.HasPrefix(s, "maerd") {
		return hasPrefix(s[5:])
	} else if strings.HasPrefix(s, "remaerd") {
		return hasPrefix(s[7:])
	} else if strings.HasPrefix(s, "esare") {
		return hasPrefix(s[5:])
	} else if strings.HasPrefix(s, "resare") {
		return hasPrefix(s[6:])
	}

	return false
}
