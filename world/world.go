package world

import (
	"fmt"

	"github.com/nmante/gomod/hello"
)

const greeting = hello.Value + " World"

// Greet is a toy example of using another submodules (github.com/nmante/gomod/hello) code
func Greet() {
	fmt.Println(greeting)
}
