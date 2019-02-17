package main

import (
	"fmt"
	"log"
	calc1 "lsbasi-go/part1/libs"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	var text string
	for {
		text = ""
		fmt.Print("calc>")
		fmt.Scanf("%s\n", &text)
		if len(text) == 0 {
			break
		}
		log.Println(text)
		interpreter := calc1.Interpreter{Text: text, Pos: 0}
		result, err := interpreter.Expr()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(result)
	}
}
