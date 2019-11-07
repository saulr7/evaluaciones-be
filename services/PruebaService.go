package services

import (
	"fmt"
)

func Say_hello() (string, error) {
	fmt.Println("It works")

	return "Hello", nil
}
