package main
#
import (
	"golangWeb/helpers"
	"log"
)

func main() {
	log.Println("hello")
	var MyVar helpers.SomeType
	MyVar.TypeName = "testing the name"
	MyVar.TypeNumber = 9
	log.Println(MyVar)
}
