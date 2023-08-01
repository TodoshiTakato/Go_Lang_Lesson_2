package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2"}

	//productList := make([]string, 0, 10) // Сделать список продуктов
	productList := map[string]int{
		"Молоко":         100,
		"Кефир":          100,
		"Йогурт":         100,
		"Кашка-Экспресс": 100,
		"Хлеб":           100,
		"Помидоры":       100,
		"Огурцы":         100,
		"Морковь":        100,
		"Мясо говяжее":   100,
		"Банан":          100}

	for command != exit {
		//fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("Введите команду\nСписок команд:\n\t%q\n\t%q\n\t%q", reg, auth, exit)

		_, err := fmt.Scan(&command)
		if err != nil {
			return
		}

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Для регистрации введите логин и пароль в формате login_password")

			_, err := fmt.Scan(&command)
			if err != nil {
				return
			}

			if slices.Contains(userList, command) { // Сделать так, что бы выводил сообщение, если пользователь уже существует
				fmt.Println("Такой пользователь уже существует!")
				fmt.Println(userList)
			} else {
				userList = append(userList, command)
				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println(message)
				fmt.Println(userList)
			}
		case auth:
			fmt.Println("Для авторизации введите логин и пароль в формате login_password")

			_, err := fmt.Scan(&command)
			if err != nil {
				return
			}

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магазин!")

					fmt.Println("Выберите интересующий вас продукт:")

					for product, count := range productList {
						fmt.Printf("%q: %d", product, count)
					}

				} else {
					fmt.Println("Вы не зарегистрированы!")
				}
			}
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
