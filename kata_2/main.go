package main

import "fmt"

// Поскольку мы не можем определить ключевые слова в Javascript
// (ну, по крайней мере, я не знаю, как это сделать),
// ваша задача - определить функцию, xor(a, b) где a и b - это два выражения,
// которые должны быть вычислены. Ваша xor функция должна вести себя так, как
// описано выше, возвращая значение true if только одно из двух выражений равно
// true, false в противном случае.

func main() {
	b := Xor(false, false)
	fmt.Println(b)

	b1 := Xor(true, false)
	fmt.Println(b1)

	b2 := Xor(false, true)
	fmt.Println(b2)

	b3 := Xor(true, true)
	fmt.Println(b3)
}

func Xor(a bool, b bool) bool {
	if a && b || !a && !b {
		return false
	}

	return true
  }