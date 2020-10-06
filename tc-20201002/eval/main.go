package main

import (
	"bufio"
	"fmt"
	"go/token"
	"go/types"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var ary [40]string = [40]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", " ", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ",", ".", "\n"}

func main() {
	start := time.Now()

	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	ans := make([]string, 1000)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		r.Scan()
		s := r.Text()

		wg.Add(1)
		func(s string, i int) {
			defer wg.Done()
			ans[i] = calc(s)
		}(s, i)
	}
	wg.Wait()

	bf := ""
	for _, v := range ans {
		if (v == " " || v == "," || v == "." || v == "\n") && bf == v {
			bf = v
			continue
		}
		fmt.Fprint(w, v)
		bf = v
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func calc(s string) string {
	var st Stack
	l := strings.Split(s, " ")
	for _, str := range l {
		if str == "+" || str == "-" || str == "*" || str == "/" {
			v1 := st.pop()
			v2 := st.pop()
			st.push(v2 + str + v1)
		} else {
			st.push(str + ".0")
		}
	}

	// eval
	expr := st.pop()
	fset := token.NewFileSet()
	tv, err := types.Eval(fset, nil, token.NoPos, expr)
	if err != nil {
		log.Fatal(err)
	}
	v, _ := strconv.ParseFloat(tv.Value.String(), 64)
	return ary[int(v)]
}

// Stack ...
type Stack []string

// pop ...
func (s *Stack) empty() bool {
	return len(*s) == 0
}

// push ...
func (s *Stack) push(i string) {
	*s = append(*s, i)
}

// pop ...
func (s *Stack) pop() string {
	index := len(*s) - 1
	res := (*s)[index]
	*s = (*s)[:index]
	return res
}
