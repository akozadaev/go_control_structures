package main

import "fmt"

func main() {
	fmt.Println("=== Форма 1: Классический цикл со счётчиком ===")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("=== Форма 2: Цикл в стиле while ===")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("=== Форма 3: Бесконечный цикл с break ===")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 5 {
			break
		}
	}
}
