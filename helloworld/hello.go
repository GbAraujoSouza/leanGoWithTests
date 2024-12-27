package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello, meu faixa "
	spanishPrefix = "Hola, meu cria "
	frenchPrefix  = "Bonjour, meu cria "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Gabriel", ""))
}
