package main

import (
	"fmt"
	"strings"
)

/*Дополните решение так, чтобы оно переворачивало все слова в переданной строке.
Слова разделяются ровно одним пробелом, без пробелов в начале или в конце.*/

func main() {
	fmt.Println(ReverseWords("The greatest victory is that which requires no battle"))
}

func ReverseWords(str string) string {
	arrStrs := strings.Split(str, " ")

	for i := 0; i < len(arrStrs) / 2; i++ {
		arrStrs[i], arrStrs[len(arrStrs) - 1 - i] = arrStrs[len(arrStrs) - 1 - i], arrStrs[i]
	}
	
	return strings.Join(arrStrs, " ")
  }