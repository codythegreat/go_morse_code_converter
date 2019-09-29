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

// raw string containing morse code
const morseCodeRaw = ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."

// define space flag
var space = flag.Bool("space", false, "adds spaces between each morse character code")

func main() {
	// check os.Args for flags, and set variables
	flag.Parse()
	// create splice out of morse characters (to cylcle through later)
	morse := strings.Split(morseCodeRaw, " ")
	returnString := ""
	// grab the last word in os.Args for conversion
	userArg := strings.ToLower(os.Args[len(os.Args)-1])
	// loop over each character, grab the byte value, use that to find morse character
	for i := range userArg {
		currentChar := fmt.Sprint(userArg[i])
		n, err := strconv.Atoi(currentChar)
		if err != nil {
			log.Fatal(err)
		}
		returnString += morse[n-97]
		if *space {
			returnString += " "
		}
	}
	fmt.Println(returnString)
}
