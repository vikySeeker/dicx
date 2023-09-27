package main

import (
	w "dicx/word"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
baic implemetation to test reading api keys from env file

	and importing packages and calling methods
*/
func main() {
	//fmt.Println("Hello World!");
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("some error: %s", err)
	}

	val := os.Getenv("THESKEY")
	fmt.Println(val)
	val = os.Getenv("DICTKEY")
	fmt.Println(val)
	fmt.Print(w.GetSelectedWord())
}
