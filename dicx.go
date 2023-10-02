package main

import (
	a "dicx/api"
	n "dicx/notify"
	w "dicx/word"
	"log"
	"os"
	"strings"
)

// variables that is used in common and used to indicate command line arguments
var include_notification bool
var word string
var err error

/*
function to determine whether the program is running in terminal mode or in background
to determine which way should the program output the result
*/
func isInTerminalMode() bool {
	env, _ := os.Stdout.Stat()
	return (env.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

/*
function to setup the the command line arguments
to define whether need to include notification in terminal mode
and whether word is give in the command line argument
*/
func setupArguments() {
	for _, args := range os.Args[1:] {
		if strings.Contains(args, "-n") {
			include_notification = true
		} else {
			word = args
		}
	}
}

/*
Main function that handle how the program flows
*/
func main() {
	setupArguments()
	if word == "" {
		word, err = w.GetSelectedWord()
		if err != nil {
			log.Fatal(err)
		}
		if err = a.GetMeaning(word); err != nil {
			log.Fatal(err)
		}

	} else {
		if err = a.GetMeaning(word); err != nil {
			log.Fatal(err)
		}
	}

	if !isInTerminalMode() {
		err = n.SendNotification(a.Result)
		if err != nil {
			log.Fatal(err)
		}
	} else if include_notification {
		err = n.SendNotification(a.Result)
		if err != nil {
			log.Fatal(err)
		}
		err = n.PrintOutput(a.Result)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err = n.PrintOutput(a.Result)
		if err != nil {
			log.Fatal(err)
		}
	}
}
