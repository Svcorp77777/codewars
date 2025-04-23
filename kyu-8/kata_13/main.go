package main

import "fmt"

/*Первый век длится с 1 года по и включительно 100 год, второй век —
с 101 года по и включительно 200 год и т. д.
Задание
По истечении года верните век, в котором он находится.*/

func main() {
	fmt.Println(century(205))
}

func century(year int) int {
	// if year % 100 == 0 {
	// 	return year / 100
	// }
	// return year / 100 + 1

	return (year + 99)/100
   }