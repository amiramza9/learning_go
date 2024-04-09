// hello.go
package main

import (
	"fmt"
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Mello, " + name + "!"
}
func main() {
	fmt.Println(Hello("amza"))
}
