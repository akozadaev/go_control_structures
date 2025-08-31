package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x больше 5")
	} else {
		fmt.Println("x не больше 5")
	}

	fmt.Println("=== Расширенная форма: if с инициализацией ===")
	if val := x * 2; val > 15 {
		fmt.Println("Удвоенное значение больше 15")
	} else {
		fmt.Println("Удвоенное значение слишком мало")
	}

	fmt.Println("===  Конструкция switch ===")
	day := "пятница"
	//day = "среда"
	switch day {
	case "понедельник":
		fmt.Println("Начало недели")
	case "пятница":
		fmt.Println("Почти выходной")
	default:
		fmt.Println("Обычный день")
	}

	switch 2 {
	case 1:
		fmt.Println("Один")
		fallthrough
	case 2:
		fmt.Println("Два") // выполнится
	case 3:
		fmt.Println("Три")
	}
	// Вывод: "Два"

	fmt.Println("===  Можно использовать switch без выражения — как замена if-else if ===")
	x = 10
	x = 0
	x = -10
	switch {
	case x < 0:
		fmt.Println("Отрицательное")
	case x == 0:
		fmt.Println("Ноль")
	default:
		fmt.Println("Положительное")
	}

	//fallthrough
}
