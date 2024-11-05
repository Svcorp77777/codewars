package main

import (
	"fmt"
)

/*Это довольно просто. Ваша задача — создать функцию, которая удаляет
первый и последний символы строки. Вам даётся один параметр — исходная
строка. Вам не нужно беспокоиться о строках, в которых меньше двух символов.*/

func main() {
	fmt.Println(RemoveChar("привет"))
}

func RemoveChar(word string) string {
	// arrWord := strings.Split(word, "")
	// return strings.Join(arrWord[1:len(arrWord)-1], "")

	var x = []rune(word)
	return string(x[1:len(x)-1])
}
