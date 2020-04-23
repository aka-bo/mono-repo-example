package main

import (
	"github.com/aka-bo/mono-repo-example/go/foo"
	"fmt"
)

const Name = "Bar"

func main() {
	fmt.Println(foo.Name, Name)
}
