package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
