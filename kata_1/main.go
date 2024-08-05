package main

import (
	"fmt"
	"strings"
)

// Учитывая 2 строки, a и b, верните строку вида short + long + short,
// с более короткой строкой снаружи и более длинной строкой внутри.
// Строки не будут одинаковой длины, но они могут быть пустыми ( zero длина ).

func main() {

	str := Solution("1", "22")

	fmt.Println(str)

}

func Solution(a, b string) string {
	var str strings.Builder
	lenA := lenStr(a)
	lenB := lenStr(b)

	if lenA > lenB {
		str.WriteString(b)
		str.WriteString(a)
		str.WriteString(b)
	} else {
		str.WriteString(a)
		str.WriteString(b)
		str.WriteString(a)
	}

	return str.String()
  }

  func lenStr(str string) int {
	var length int

	for range str {
		length ++
	}

	return length
  }
