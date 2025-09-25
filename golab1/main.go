package main

import (
	"fmt"
	"time"
)

func main() {
	//получаение даты и времени с помощью import time
	now := time.Now()

	fmt.Println("----------------------------------------------------")
	//вывод всей информации
	fmt.Println("Дата и время: ", now)

	var a int = 1
	var b float64 = 2.2
	var c string = "три"
	var d bool = true

	fmt.Println("----------------------------------------------------")

	fmt.Println("int", a)
	fmt.Println("float64", b)
	fmt.Println("string", c)
	fmt.Println("bool", d)

	fmt.Println("----------------------------------------------------")

	z := 1
	x := 2.2
	s := "строка"
	v := false

	fmt.Println("int", z)
	fmt.Println("float64", x)
	fmt.Println("string", s)
	fmt.Println("bool", v)

	fmt.Println("----------------------------------------------------")

	result := a + z
	fmt.Println("Результат сложения ", result)

	fmt.Println("----------------------------------------------------")

	fmt.Println("Арифметичесике операции над 2 целыми числами")

	var first, second int

	fmt.Print("Введите первое целое число: ")
	fmt.Scan(&first)

	fmt.Print("Введите второе целое число: ")
	fmt.Scan(&second)

	sum := first + second
	diff := first - second
	product := first * second
	quotient := first / second
	remainder := first % second

	fmt.Printf("%d + %d = %d\n", first, second, sum)
	fmt.Printf("%d - %d = %d\n", first, second, diff)
	fmt.Printf("%d * %d = %d\n", first, second, product)
	fmt.Printf("%d / %d = %d\n", first, second, quotient)
	fmt.Printf("%d %% %d = %d\n", first, second, remainder)

	fmt.Println("----------------------------------------------------")

	fmt.Println("Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой")
	fmt.Println(calculate(6.3, 3))
	fmt.Println("----------------------------------------------------")
	fmt.Println("Среднее значение 3 чисел")
	fmt.Println(avg(3, 2, 1))
}

func calculate(one, two float64) (float64, float64) {
	sum := one + two
	difference := one - two
	return sum, difference
}

func avg(one, two, three int) int {
	result := (one + two + three) / 3
	return result
}
