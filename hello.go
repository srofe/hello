package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishPrefix + name
	}
	if language == french {
		return frenchPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
