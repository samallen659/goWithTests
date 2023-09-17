package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Sam"))
}

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return fmt.Sprintf("%s%s", englishHelloPrefix, name)
}
