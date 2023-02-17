package main

import (
	"fmt"
	"math"
	"sort"
)

func task1() {
	var s, v, v1, t float64

	fmt.Scan(&s)
	fmt.Scan(&v)
	fmt.Scan(&v1)
	t = s / (v + v1)

	fmt.Println(t)
}

func task2() {
	s := []int{4, 2, 3, 6, 7, 4, 2, 5, 7, 4, 3, 6, 8, 2, 2, 2, 2, 2}
	sort.Ints(s)
	fmt.Println(s)
}

func task3() {
	var day, hours, min, sec float64

	fmt.Scan(&sec)
	day = sec / 86400
	hours = sec / 3600
	min = sec / 60

	println(day, hours, min, sec)
}

func task4() {
	var year int = 3

	if year%4 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func task5() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	d := math.Pow(b, 2) - 4*a*c
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)

	fmt.Println("Roots:", x1, x2)
}

func task6() {
	a := 1.0
	b := 5.0
	c := 6.0
	sum := -b / a
	product := c / a

	x1 := sum/2.0 + (sum/2.0)*(sum/2.0) - product
	x2 := sum/2.0 - (sum/2.0)*(sum/2.0) - product

	fmt.Print("Корни уравнения", a, b, c)
	fmt.Print(x1)
	fmt.Print(x2)
}

func task7() {
	a := 1
	b := 2
	c := 3
	d := 2
	count := 4
	if a == b {
		count--
	}
	if a == c || b == c {
		count--
	}
	if a == d || b == d || c == d {
		count--
	}
	fmt.Println(count)
}

func main() {
	//task1()
	//task2()
	//task3()
	//task4()
	//task5()
	//task6()
	task7()
}
