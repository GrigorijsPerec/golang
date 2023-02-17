package main

import (
	"fmt"
	"math"
)

func tas1() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var s int = a * b
	var p int = a + b + a + b
	fmt.Println("S - ", s)
	fmt.Println("P - ", p)
}

func tas2() {
	var a, b, c int
	var v, saa int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	v = a * b * c
	saa = 2 * (a*b + a*c + b*c)

	fmt.Println(v)
	fmt.Println(saa)
}
func tas3() {
	var km, h float32
	var ms float32
	var mmin float32

	fmt.Scan(&km)
	fmt.Scan(&h)

	ms = (km / 1000) / (h / 3600)
	mmin = km * 16.6

	fmt.Println(ms)
	fmt.Println(mmin)
}

func tas4() {
	var v, t, s int

	fmt.Scan(&v)
	fmt.Scan(&t)
	fmt.Scan(&s)

	s = v * t

	fmt.Println(s)
}

func tas5() {
	var c, r, p, s float64
	fmt.Scan(&r)
	p = 3.14

	c = 2 * p * r
	s = p * math.Pow(r, 2)

	fmt.Println(s)
	fmt.Println(c)
}

func tas6() {
	var n, s, h float64
	fmt.Scan(&n)
	fmt.Scan(&s)
	h = math.Pow(n, s)
	fmt.Println(h)
}

func main() {
	//tas1()
	//tas2()
	//tas3()
	//tas4()
	//tas5()
	tas6()

}
