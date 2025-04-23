package main

import (
	"fmt"
	"strings"
)

// Напишите функцию для преобразования имени в инициалы. Это ката состоит строго из двух слов с одним пробелом между ними.
// Выходные данные должны состоять из двух заглавных букв с точкой, разделяющей их.
// Это должно выглядеть примерно так:
// Sam Harris => S.H
// patrick feeney => P.F

func main() {
	str := AbbrevName("Иван Петров")
	fmt.Println(str)
}

func AbbrevName(ns string) string {
	var res strings.Builder

	names := strings.Split(ns, " ")

	for _, val := range names[0] {
		res.WriteString(strings.ToUpper(string(val)))

		break
	}

	res.WriteString(".")

	for _, val := range names[1] {
		res.WriteString(strings.ToUpper(string(val)))

		break
	}

	return res.String()
}
