package main

import "fmt"

const englishHelloPrefix = "Hellow, "

func Hello(name string) string {
	if name == "" {
		name = "Worldy"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
