package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// In Go, if a function starts with a lowercase, it means it's private.
// Capital case means they are public.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	// Just `return` here is equivalent to `return prefix`
	return
}

func main() {
	fmt.Println(Hello("Victor", "English"))
}
