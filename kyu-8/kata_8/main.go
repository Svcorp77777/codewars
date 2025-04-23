package main

import "fmt"

// Тимми и Сара думают, что они влюблены, но в том месте, где они живут,
// они узнают об этом, только когда сорвут по цветку. Если у одного из цветов
// чётное количество лепестков, а у другого — нечётное, значит, они влюблены.

// Напишите функцию, которая будет принимать количество лепестков у каждого
// цветка и возвращать значение true, если они влюблены, и false, если нет.

func main() {
	fmt.Println(LoveFunc(1, 4))
	fmt.Println(LoveFunc(2, 2))
	fmt.Println(LoveFunc(0, 1))
	fmt.Println(LoveFunc(0, 0))
}

func LoveFunc(flower1, flower2 int) bool {
	return flower1 % 2 != flower2 % 2
}
