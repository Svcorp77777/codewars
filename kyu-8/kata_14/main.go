package main

import "fmt"

/*Очень просто: дано число (целое / десятичное / оба варианта в зависимости от языка),
найдите противоположное ему (аддитивно обратное).
Примеры:
1: -1
14: -14
-34: 34*/

func main() {
	fmt.Println(Opposite(1))
	fmt.Println(Opposite(14))
	fmt.Println(Opposite(-34))
	fmt.Println(Opposite(0))
}

func Opposite(value int) int {
	return -value 
}