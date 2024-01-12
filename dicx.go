package main

import (
	a "dicx/api"
	n "dicx/notify"
	w "dicx/word"
	"fmt"
	"log"
	"os"
	"strings"
)

// variables that is used in common and used to indicate command line arguments
var include_notification bool
var word string
var err error
var help_flag bool

/*
function that prints the help section flag: -h
*/

func printHelp() {
	fmt.Println("Usage: dicx [options] <word>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("\t -h\t\t displays help section")
	fmt.Println("\t -n\t\t force send notification in terminal mode")
	fmt.Println()
	fmt.Println("Working:")
	fmt.Println("\tdicx is simple command line tool that shows output for a word. dicx can be used in two modes \n\t\t1) Shortcut mode \n\t\t2) Terminal mode")
	fmt.Println()
	fmt.Println("\tShortcut mode:\n\tTo use it in shortcut mode bind the command `dicx` with any keyboard shortcut for eg. bind command `dicx` with ctrl+m key.")
	fmt.Println("\tTerminal mode:\n\tTo use it in terminal mode run the command dicx with a word as a argument.")
}

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
	if len(os.Args) < 2 {
		help_flag = true
		return
	}
	for _, args := range os.Args[1:] {
		if strings.Contains(args, "-h") {
			help_flag = true
			break
		}

		if strings.Contains(args, "-n") {
			include_notification = true
			continue
		}
		word = args
	}
}

/*
Main function that handle how the program flows
*/
func main() {
	setupArguments()
	if help_flag {
		printHelp()
		os.Exit(0)
	}
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

	if !isInTerminalMode() || include_notification {
		n.SendNotification(a.Result)
	} else {
		n.PrintOutput(a.Result)
	}
}
