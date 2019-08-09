package hangmangame

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var text string
var scanner = bufio.NewScanner(os.Stdin) //this will enable us to read user input from console
var parts int
var phrase string

func Playhangman() {
	phrase = "HAPPY"
	playgame(phrase, &parts)
}

func hangmangame(phrase string, text string) {
	if phrase != "" && parts < 7 {
		if strings.Contains(phrase, text) {
			fmt.Println("You are correct, continue playing the game")
			phrase = strings.Replace(phrase, text, "", 1)
			playgame(phrase, &parts)
		} else {
			fmt.Println("Oops! That letter is not in the phrase I chose. Better luck next time", text)
			_ = hangman(&parts)
			playgame(phrase, &parts)
		}
	}
}

func playgame(phrase string, parts *int) (){
	if phrase == "" {
		fmt.Println("Congratualtions!!!! You guessed all the letters correctly")
	} else if *parts > 5 {
		fmt.Println("Sorry, hangman is fully hanging now")
	} else {
		fmt.Println("Guess a letter!")
	scanner.Scan()                            //Scan is one of the function amongst many in bufio package
	text = scanner.Text()
	hangmangame(phrase, text)
	}
}

func hangman(parts *int) (*int){
	*parts++
	r := *parts
	return &r
}