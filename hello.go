package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	targetGreeting := getHelloByLang(lang)

	if name == "" {
		return targetGreeting + "World" + "!"
	}
	return targetGreeting + name + "!"
}

func getHelloByLang(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
