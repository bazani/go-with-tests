package main

import "fmt"

func Hello(name string) string {
	return "Hellow, " + name
}

func main() {
	fmt.Println(Hello("wordly"))
}
