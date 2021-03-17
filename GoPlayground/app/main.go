package main

import (
	. "GoPlayground/app/joker"
	"fmt"
)

func main() {
	const url = "https://api.chucknorris.io/jokes/random"
	joker := NewJoker(url)
	joke, code := joker.MakeJoke()
	if code != 200 {
		fmt.Println(code)
	} else {
		fmt.Println(joke)
	}
}
