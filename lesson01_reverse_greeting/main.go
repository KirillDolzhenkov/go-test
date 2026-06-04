package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	// Переворачиваем строку "Hello, OTUS!"
	result := reverse.String("Hello, OTUS!")
	fmt.Println(result)
}
