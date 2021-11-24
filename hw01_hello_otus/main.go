package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {

	hl := "Hello, OTUS!"
	result := stringutil.Reverse(hl)
	fmt.Println(result)
}
