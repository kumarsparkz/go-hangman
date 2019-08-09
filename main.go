package main

import (
	// import default packages from go
	"fmt"
	// importing packages from other libraries
	"github.com/kumarsparkz/Gohangman/hangmangame"
)

// you can call only main function in go and that should call any other internal functions
func main(){
	fmt.Println("Welcome to Kumarsparkz Golang Library")
	hangmangame.Playhangman()
}
