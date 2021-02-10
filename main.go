package main

import (
	"auto-garden/argparsing"
	"fmt"
)


func main() {

	a := argparsing.ParseArgs()
	fmt.Printf("%+v", *(a.DbPassword))
}
