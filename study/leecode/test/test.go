package main

import (
	"fmt"
	"strconv"
)

//Stack 栈
type Stack struct {
	values []string
}

func (s *Stack) poll() (v string) {
	if len(s.values) == 0 {
		panic("stack is nil")
	}
	v = s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return v
}

//HighStack 高位运算栈
type HighStack struct {
	Stack
}

func (s *HighStack) result() string {
	lowStack := &LowStack{}
	split := s.values
	for i := 0; i < len(split); i++ {
		v := split[i]
		lowStack.add(v)
	}
	return lowStack.values[0]
}
func (s *HighStack) add(v string) {
	if !isNumber(v) && v != ")" {
		s.values = append(s.values, v)
		return
	}
	if v == ")" {
		temp := s.poll()
		lowStack := &LowStack{}
		for ; temp != "("; temp = s.poll() {
			lowStack.add(temp)
		}
		v = lowStack.poll()
	}
	if len(s.values) > 1 {
		symbol := s.values[len(s.values)-1]
		if symbol == "*" || symbol == "/" {
			s.poll()
			v = calculation(symbol, s.poll(), v)
		}
	}
	s.values = append(s.values, v)
}

//LowStack 低位运算栈
type LowStack struct {
	Stack
}

func (s *LowStack) add(v string) {
	if isNumber(v) {
		if len(s.values) > 1 {
			symbol := s.values[len(s.values)-1]
			if symbol == "+" || symbol == "-" {
				s.poll()
				v = calculation(symbol, s.poll(), v)
			}
		}
	}
	s.values = append(s.values, v)
}
func (s *LowStack) poll() (v string) {
	if len(s.values) == 0 {
		panic("stack is nil")
	}
	v = s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return v
}

func isNumber(v string) bool {
	if _, e := strconv.Atoi(v); e != nil {
		return false
	}
	return true
}
func calculation(symbol, v1, v2 string) (v string) {
	pre, _ := strconv.Atoi(v1)
	pos, _ := strconv.Atoi(v2)
	switch symbol {
	case "+":
		v = strconv.Itoa(pre + pos)
	case "-":
		v = strconv.Itoa(pre - pos)
	case "*":
		v = strconv.Itoa(pre * pos)
	case "/":
		v = strconv.Itoa(pre / pos)
	}
	return v
}

func main() {
	stack := &HighStack{}
	split := []string{}
	str := "1*5*(3+4*6)-2"
	tempStr := ""
	for i, v := range str {
		temp := string(v)
		if temp == "+" || temp == "-" || temp == "*" || temp == "/" || temp == "(" || temp == ")" {
			if len(tempStr) > 0 {
				split = append(split, tempStr)
			}
			split = append(split, temp)
			tempStr = ""
		} else {
			tempStr += temp
		}
		if i == len(str)-1 {
			split = append(split, tempStr)
		}
	}
	fmt.Println(split)
	for i := 0; i < len(split); i++ {
		v := split[i]
		stack.add(v)
	}
	fmt.Println(stack.result())
}
