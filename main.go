/*
package main

import (
	"fmt"
	"log"
)

func main() {
	n := 0
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %d\n", n)
}
*/

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

	fmt.Printf("Вы ввели: %s\n", input)
}
