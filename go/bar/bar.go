package main

import (
	"github.com/aka-bo/mono-repo-example/go/foo"
	"fmt"
)

const Greeting = "Hello"
const Name = "Bar"

func main() {
	fmt.Println(Greeting)
	fmt.Println(foo.Name, Name)
}
