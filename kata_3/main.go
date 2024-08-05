package main

import "fmt"

// Напишите программу, в которой Алекс может ввести (n), сколько раз крутится обруч, 
// и она вернет ему ободряющее сообщение:
// Если Алекс получит 10 или более обручей, верните строку "Great, now move on to tricks".
// Если он не получит 10 обручей, верните строку "Keep at it until you get it".

func main() {
	n := HoopCount(3)
	fmt.Println(n)

	n1 := HoopCount(12)
	fmt.Println(n1)
}

func HoopCount(n int) string {
	if n < 10 {
		return "Keep at it until you get it"
	}

	return "Great, now move on to tricks"
}
