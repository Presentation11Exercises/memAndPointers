package main

import (
	"fmt"
	)

func main() {
  myValue := "whatsup"
  fmt.Println(*&myValue)
}
