package main

import "fmt"

// Создайте функцию, которая возвращает массив целых чисел от n до 1, где n>0.
// Пример : n=5 --> [5,4,3,2,1]

func main() {
	fmt.Println(ReverseSeq(29))
}

func ReverseSeq(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = n - i
	}
	return arr
}
