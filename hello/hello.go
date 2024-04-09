// hello.go
package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Mello, "+name+"!"
}
func main() {
	fmt.Println(Hello("amza"))
}
