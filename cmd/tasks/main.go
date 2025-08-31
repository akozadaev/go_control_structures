package main

import "fmt"

func main() {
	//Классификация чисел в диапазоне
	сlassificationNumbersInRange()
	//Задача. Поиск первого простого числа в диапазоне
	primeNumber()
	//Симуляция авторизации с попытками
	simulateAuthorizationAttempts()
	//	Поиск "идеального" числа
	perfectNumber()
}

func сlassificationNumbersInRange() {
	//Задание:
	//	Для чисел от 1 до 20 определите тип числа с помощью switch:
	//
	//		"чётное", если делится на 2
	//		"нечётное", если не делится
	//		"особое", если это 1 или 2
	//		Используйте switch без выражения (как if-else цепочку).

	for i := 1; i <= 20; i++ {
		switch {
		case i == 1 || i == 2:
			fmt.Printf("%d — особое\n", i)
		case i%2 == 0:
			fmt.Printf("%d — чётное\n", i)
		default:
			fmt.Printf("%d — нечётное\n", i)
		}
	}
}

func primeNumber() {
	//Задание:
	//	Напишите программу, которая находит первое простое число в диапазоне от start до end (включительно). Если простое число не найдено, выведите 0.
	//	Используйте вложенные циклы и break с меткой для оптимизации.
	start, end := 10, 25

	var result int = 0
	found := false

search:
	for i := start; i <= end; i++ {
		if i < 2 {
			continue
		}
		if i == 2 {
			result = 2
			found = true
			break
		}
		if i%2 == 0 {
			continue
		}

		for j := 3; j*j <= i; j += 2 {
			if i%j == 0 {
				continue search
			}
		}
		result = i
		found = true
		break
	}

	if found {
		fmt.Printf("Первое простое число: %d\n", result)
	} else {
		fmt.Println("Простое число не найдено")
	}

}

func simulateAuthorizationAttempts() {
	//Задание:
	//	Напишите программу, которая позволяет пользователю ввести пароль до 3 раз.
	//		Правильный пароль — "secret123".
	//		Используйте цикл for в стиле while, счётчик попыток и if для проверки.

	const correct = "secret123"
	var input string
	attempts := 0
	maxAttempts := 3

	for attempts < maxAttempts {
		fmt.Printf("Введите пароль (осталось %d): ", maxAttempts-attempts)
		fmt.Scanln(&input)

		if input == correct {
			fmt.Println("Доступ разрешён")
			return
		}

		attempts++
	}

	fmt.Println("Доступ запрещён: превышено количество попыток")
}

func perfectNumber() {
	//Задание:
	//	Найдите первое "идеальное число" в диапазоне от 2 до 30.
	//	Идеальное число — это число, равное сумме своих делителей (кроме себя).
	//	Например, 6 = 1 + 2 + 3.
	//	Если не найдено — выведите 0.

	var perfect int = 0

search:
	for n := 2; n <= 30; n++ {
		sum := 0
		for d := 1; d < n; d++ {
			if n%d == 0 {
				sum += d
			}
		}
		if sum == n {
			perfect = n
			break search
		}
	}

	if perfect != 0 {
		fmt.Printf("Первое идеальное число: %d\n", perfect)
	} else {
		fmt.Println("Идеальное число не найдено")
	}
}
