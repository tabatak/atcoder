package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var ary [40]string = [40]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", " ", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ",", ".", "\n"}
var order map[string]int = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
var ans [1000]string

func main() {
	start := time.Now()

	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

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
			st.push(str)
		}
	}

	// 逆ポーランド記法に戻す
	siki := st.pop()
	var stP Stack
	var porlandSiki Queue
	for _, v := range siki {
		str := string(v)
		if str == "+" || str == "-" || str == "*" || str == "/" {
			if stP.empty() {
				stP.push(str)
			} else {
				for !stP.empty() && order[stP.peek()] >= order[str] {
					a := stP.pop()
					porlandSiki.push(a)
				}
				stP.push(str)
			}
		} else {
			porlandSiki.push(string(v))
		}
	}
	for !stP.empty() {
		porlandSiki.push(stP.pop())
	}

	return ary[calcRPorland(strings.Join(porlandSiki, ""))]
}

func calcRPorland(s string) int {
	var st NumStack
	for _, v := range s {
		str := string(v)
		if str == "+" {
			st.push(st.pop() + st.pop())
		} else if str == "-" {
			v1 := st.pop()
			v2 := st.pop()
			st.push(v2 - v1)
		} else if str == "*" {
			st.push(st.pop() * st.pop())
		} else if str == "/" {
			v1 := st.pop()
			v2 := st.pop()
			st.push(v2 / v1)
		} else {
			v, _ := strconv.ParseFloat(str, 64)
			st.push(v)
		}
	}
	return int(st.pop())
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
	if s.empty() {
		return ""
	} else {
		index := len(*s) - 1
		res := (*s)[index]
		*s = (*s)[:index]
		return res
	}
}

// NumStack ...
type NumStack []float64

// pop ...
func (s *NumStack) empty() bool {
	return len(*s) == 0
}

// push ...
func (s *NumStack) push(i float64) {
	*s = append(*s, i)
}

// pop ...
func (s *NumStack) pop() float64 {
	if s.empty() {
		return 0
	} else {
		index := len(*s) - 1
		res := (*s)[index]
		*s = (*s)[:index]
		return res
	}
}

// pop ...
func (s *Stack) peek() string {
	if s.empty() {
		return ""
	} else {
		index := len(*s) - 1
		res := (*s)[index]
		return res
	}
}

// Queue ...
type Queue []string

// pop ...
func (q *Queue) empty() bool {
	return len(*q) == 0
}

// push ...
func (q *Queue) push(i string) {
	*q = append(*q, i)
}

// pop ...
func (q *Queue) pop() string {
	if q.empty() {
		return ""
	} else {
		res := (*q)[0]
		*q = (*q)[1:]
		return res
	}
}
