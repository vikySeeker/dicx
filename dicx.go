package main

import (
	a "dicx/api"
	w "dicx/word"
	"fmt"
)

/*
baic implemetation to test reading api keys from env file

	and importing packages and calling methods
*/
func main() {
	res, err := w.GetSelectedWord()
	if err != 0 {
		fmt.Println("some error occured: " + res)
	}
	fmt.Println("Selected word is: " + res)
	a.GetMeaning(res)
}
