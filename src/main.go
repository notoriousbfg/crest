package main

import (
	"crest/lexer"
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("./basic.crest")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	l := lexer.New(input)
	fmt.Printf("%+v", l.Tokens)
	fmt.Println()
}
