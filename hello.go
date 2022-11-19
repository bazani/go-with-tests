package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hellow, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portuguesePrefix = "Ola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "Worldy"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portuguesePrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
