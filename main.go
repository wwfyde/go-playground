package main

import (
	"fmt"
	"github.com/wwfyde/go-playground/example"
	"github.com/wwfyde/go-playground/files/read_file"
)

func main() {
	example.Hello()
	fmt.Println("hello, gin!")
	fmt.Println(read_file.Read())
}
