package main

import (
	a "dicx/api"
	n "dicx/notify"
	w "dicx/word"
	"fmt"
	"strings"
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
	res = strings.TrimSpace(res)
	res = strings.Split(res, " ")[0]
	res = strings.ReplaceAll(res, ".", "")
	fmt.Println("Selected word is: " + res)
	if a.GetMeaning(res) == 1 {
		n.SendNotification(a.Result)
	}
}
