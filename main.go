package main

import (
	"fmt"
	"log"
)

func main() {
	var input string

	fmt.Print("Введите данные: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Вы ввели данные : %s\n", input)
}
