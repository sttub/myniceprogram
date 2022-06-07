package main

import (
	"github.com/sttub/myniceprogram/helpers"
	"log"
)

func main() {
	log.Println("hello")

	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
