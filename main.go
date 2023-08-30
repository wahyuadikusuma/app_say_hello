package main

import (
	"fmt"

	say_hello "github.com/wahyuadikusuma/go-modules"
	say_hello_2 "github.com/wahyuadikusuma/go-modules/v2"
)

func main() {
	fmt.Println(say_hello.SayHello("Wahyu"))
	fmt.Println(say_hello_2.SayHello("Wahyu", "Solo"))
}