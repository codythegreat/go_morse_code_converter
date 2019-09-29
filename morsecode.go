// a simple text to morse code converter for the command line
package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
)

// split raw string containing morse code by spaces
var morse []string = strings.Split(".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --..", " ")

// define space flag
var space = flag.Bool("space", false, "adds spaces between each morse character code")

func main() {
	// check os.Args for flags, and set variables
	flag.Parse()
	// grab words is os.Args for conversion
	userArg := getWords()
	// loop over each character, grab the byte value, use that to find morse character
	fmt.Println(convChars(userArg))
}

func getWords() string {
	words := ""
	if *space {
		words = strings.ToLower(strings.Join(os.Args[2:len(os.Args)], ""))
	} else {
		words = strings.ToLower(strings.Join(os.Args[1:len(os.Args)], ""))
	}
	return words
}

func convChars(letters string) string {
	convertedChars := ""
	for i := range letters {
		n, err := strconv.Atoi(fmt.Sprint(letters[i]))
		if err != nil {
			log.Fatal(err)
		}
		convertedChars += morse[n-97]
		if *space {
			convertedChars += " "
		}
	}
	return convertedChars
}
