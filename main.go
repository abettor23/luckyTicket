package main

import (
	"fmt"
)

func main() {
	fmt.Println("Узнаем, какой билет вам достался - зеркальный, счастливый или обычный!")

	fmt.Println("Введите четырёхзначный номер билета:")
	var ticketNum int
	fmt.Scan(&ticketNum)
	if ticketNum < 1000 || ticketNum > 9999 {
		fmt.Println("Ошибка! Номер билета неверен.")
		return
	}

	num1 := ticketNum / 1000
	num2 := (ticketNum / 100) % 10
	num3 := (ticketNum / 10) % 10
	num4 := ticketNum % 10

	if num1 == num4 && num2 == num3 && num1+num2 == num3+num4 {
		fmt.Println("Вам достался и счастливый и зеркальный билет!")
	} else if num1 == num4 && num2 == num3 {
		fmt.Println("Вам достался зеркальный билет!")
	} else if num1+num2 == num3+num4 {
		fmt.Println("Вам достался счастливый билет!")
	} else {
		fmt.Println("Вам достался обычный билет!")
	}
}
