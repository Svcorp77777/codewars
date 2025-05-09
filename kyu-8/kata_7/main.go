package main

import (
	"fmt"
)

// Вы отправились в поход с друзьями далеко от дома, но когда пришло время
// возвращаться, вы поняли, что у вас заканчивается топливо, а ближайшая заправка
// находится в 50 милях от вас! Вы знаете, что в среднем ваш автомобиль расходует
// около 25 миль на галлон. У вас осталось 2 галлонов.

// Учитывая эти факторы, напишите функцию, которая сообщит вам, можно ли добраться до насоса.

// Функция должна возвращать, true если это возможно, и false если нет.

func main() {
	fmt.Println(ZeroFuel(50, 25, 2))
}

func ZeroFuel(distanceToPump int, mpg int, fuelLeft int) bool {
	res := float64(distanceToPump) / (float64(fuelLeft) * float64(mpg))
	return res <= 1.00
}
