package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var THESAPI = "https://www.dictionaryapi.com/api/v3/references/thesaurus/json/"
var DICTAPI = "https://www.dictionaryapi.com/api/v3/references/thesaurus/json/"

// prepares the url responsible to fetch the data from api
func PrepareUrl(word string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("some error: %s", err)
	}
	key := os.Getenv("THESKEY")
	THESAPI = THESAPI + word + "?key=" + key
}

// makes http request to the prepared url and extracts the meaning, definition and example
func GetMeaning(word string) {
	PrepareUrl(word)
	//fmt.Println(THESAPI)
	response, err := http.Get(THESAPI)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := string(responseData)
	fmt.Print(data)

}
