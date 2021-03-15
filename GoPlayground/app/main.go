package main

import (
	. "GoPlayground/app/printer"
	"fmt"
)

func main() {
	const content = "SOME_CONTENT"
	printer := NewPrinter("lol")
	fmt.Println(printer.Value())
}
