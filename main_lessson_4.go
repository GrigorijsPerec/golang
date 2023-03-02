package main

import (
	"fmt"
	"math"
)

func task1() {

	a := 0
	b := 1

	for b <= 1000000 {
		fmt.Println(b)

		a, b = b, a+b

	}
}

func main() {
	var N, x, y int
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	max := int(math.Sqrt(float64(N))) 
	found := false                    

	for x = 1; x <= max; x++ {
		for y = 1; y <= max; y++ {
			if x*x+y*y == N {
				fmt.Printf(x, y)
				found = true
			}
		}
	}

	if !found {
		fmt.Println("Решений нет")
	}
}

func task2() {
	var n, reversed int
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&n)

	for n > 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	fmt.Printf("Перевернутое число: %d\n", reversed)
}

func main() {
	task1()
	task2()
	task3()
}
