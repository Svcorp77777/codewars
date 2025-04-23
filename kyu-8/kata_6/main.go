package main

import (
	"fmt"
	"strings"
)

// Напишите функцию для разделения строки на слова и преобразования
// её в массив слов.
// Пример:
// "Robin Singh" ==> ["Robin", "Singh"]
// "I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

func main() {
	str1 := "Robin Singh"
	res1 := StringToArray(str1)
	fmt.Println(res1)

	str2 := "I love arrays they are my favorite"
	res2 := StringToArray(str2)
	 fmt.Println(res2)
}

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}