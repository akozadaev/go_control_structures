package main

import "fmt"

func main() {
	fmt.Println("=== Выход Только из внутреннего ===")
	breakInner()
	fmt.Println("=== Выход из внешнего ===")
	breakOuter()
}

func breakInner() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Внешний цикл: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("  Внутренний цикл: %d\n", j)
			if i == 1 && j == 1 {
				fmt.Printf("    Глубоко-глубоко: %d, %d\n", i, j)
				break
			}
		}
	}
	fmt.Println("Вот и всё")
}

func breakOuter() {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Printf("Внешний цикл: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("  Внутренний цикл: %d\n", j)
			if i == 1 && j == 1 {
				fmt.Printf("    Глубоко-глубоко: %d, %d\n", i, j)
				break outerLoop
			}
		}
	}
	fmt.Println("Вот и всё")
}
