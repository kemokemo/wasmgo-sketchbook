package main

import (
	"fmt"

	"github.com/gopherjs/gopherwasm/js"
)

func pressed(args []js.Value) {
	fmt.Println("button clicked")
}

func registerCallbacks() {
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", js.NewCallback(pressed))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Go wasm initialized.")
	registerCallbacks()

	<-c
}
