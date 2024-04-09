// hello.go
package main

import (
	"fmt"
)

const engPrefix = "Hello, "

var prefix = engPrefix

const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, lang string) string {

	if name == "" {
		name = "World"
	}
	return greeter(name, lang)

}
func greeter(name string, lang string) (prefix string) {
	switch lang {
	// if lang == spanish {
	// 	prefix = spannishPrefix
	// }
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix

	}
	return prefix + name + "!"
}
func main() {
	fmt.Println(Hello("Joma", "Spanish"))
}
