package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, i)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		var runeVal rune
		runeVal, w = utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeVal, i)

		examineRune(runeVal)
	}

}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
	
}
