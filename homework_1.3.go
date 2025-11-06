package main

import "fmt"

// Написать программу на ЯП Go. Условие: На вход программе подаются 5 натуральных чисел.
// Задача программы - считать эти числа, отсортировать по убыванию и вывести результат сортировки на экран.
// Также вывести на экран другие параметры считанной последовательности в формате:
//  "Самое большое число: {число}", "Самое маленькое число: {число}", "Среднее арифметическое: {число}"

// Пример работы программы:
// выход: 5 2 4 3 1
// выход:
// "Отсортированные элементы: 5 4 3 2 1
// Самое большое число: 5
// Самое маленькое число: 1
// Среднее-арифметическое: 3

func main() {
	operationsWithNumbers()
}

func operationsWithNumbers(){
	var num1, num2, num3, num4, num5, average int

	fmt.Println("Введите первое число для сравнения")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число для сравнения")
	fmt.Scan(&num2)
	fmt.Println("Введите третье число для сравнения")
	fmt.Scan(&num3)
	fmt.Println("Введите четвертое число для сравнения")
	fmt.Scan(&num4)
	fmt.Println("Введите пятое число для сравнения")
	fmt.Scan(&num5)

	for i:=0;i<5;i++{
		if num1>num2 {
			num1,num2 = num2,num1
		}
		if num2>num3 {
			num2,num3 =num3,num2
		}
		if num3>num4 {
			num3,num4 =num4,num3
		}
		if num4>num5 {
			num4,num5 =num5,num4
		}
	}
	fmt.Printf("Все числа по возрастанию: %d, %d, %d, %d, %d \n", num1, num2, num3, num4, num5)
	average = (num1 + num2 + num3 + num4 + num5) / 5
	fmt.Printf("Средне арифметическое: %d \n", average)
	fmt.Printf("Минимальное число: %d \n", num1)
	fmt.Printf("Максимальное число: %d \n", num5)
}