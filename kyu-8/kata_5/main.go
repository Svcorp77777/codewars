package main

import (
	"fmt"
	"strings"
)

// Если не можешь уснуть, просто считай овец!!
// Задание:
// Если задано неотрицательное целое число, 3 например, верни строку с ропотом: 
// "1 sheep...2 sheep...3 sheep...". Входные данные всегда будут действительными,
// т. Е. Без отрицательных целых чисел.

func main() {
	str := countSheep(0)
	fmt.Println(str)

}

func countSheep(num int) string {
	var res strings.Builder
	for i := 1; i <= num; i++ {
		res.WriteString(fmt.Sprintf("%d sheep...", i))
	}

	return res.String()
}
