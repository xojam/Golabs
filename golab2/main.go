package main

import "fmt"

func checkNumberSign(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func stringLength(s string) int {
	return len(s)
}

func average(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	fmt.Println("------------------Проверка четности числа-------------------")
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	fmt.Println("Проверка на чётность")
	if num%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
	fmt.Println("--------------------------Функция проверки знака числа------------------------------")

	fmt.Println(checkNumberSign(5))  // Positive
	fmt.Println(checkNumberSign(-3)) // Negative
	fmt.Println(checkNumberSign(0))  // Zero

	fmt.Println("---------------------Вывод чисел от 1 до 10------------------------")
	for i := 0; i <= 10; i++ {
		fmt.Println(i, "")
	}

	fmt.Println("-----------------------Функция для определения длины строки-------------------------")

	fmt.Println(stringLength("Hello"))
	fmt.Println(stringLength("GoLang"))
	fmt.Println(stringLength(""))

	fmt.Println("-----------------------Структура Rectangle с методом площади-------------------------")

	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("Площадь прямоугольника: %.2f\n", rect.Area())

	fmt.Println("-------------------------Функция вычисления среднего значения-----------------------")

	fmt.Printf("Среднее 4 и 6: %.1f\n", average(4, 6))
	fmt.Printf("Среднее 10 и 20: %.1f\n", average(10, 20))
	fmt.Printf("Среднее 7 и 3: %.1f\n", average(7, 3))

	fmt.Println("--------------------------------------------------------")

}

type Rectangle struct {
	Width  float64
	Height float64
}
