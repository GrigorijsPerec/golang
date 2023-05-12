package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("orders.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	orders := make(map[string]int)

	customersByProduct := make(map[string][]string)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) != 2 {
			continue
		}
		customer := words[0]
		product := words[1]

		orders[customer]++

		customersByProduct[product] = append(customersByProduct[product], customer)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Всего покупателей: %d\n", len(orders))
	fmt.Printf("Всего заказов: %d\n", len(customersByProduct))

	fmt.Println("Количество заказов для каждого покупателя:")
	for customer, count := range orders {
		fmt.Printf("%s: %d\n", customer, count)
	}

	maxOrders := 0
	maxOrderCustomers := make([]string, 0)

	for customer, count := range orders {
		if count > maxOrders {
			maxOrders = count
			maxOrderCustomers = []string{customer}
		} else if count == maxOrders {
			maxOrderCustomers = append(maxOrderCustomers, customer)
		}
	}

	fmt.Printf("Максимальное количество заказов: %d\n", maxOrders)
	fmt.Printf("Количество покупателей с максимальным количеством заказов: %d\n", len(maxOrderCustomers))
	fmt.Println("Покупатели с максимальным количеством заказов:")
	for _, customer := range maxOrderCustomers {
		fmt.Println(customer)
	}

	product := "яблоко"
	fmt.Printf("Покупатели, заказавшие товар %s: %v\n", product, customersByProduct[product])
}
