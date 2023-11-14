package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// json data struct
type apiSuccessData []struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text  string `json:"text"`
		Audio string `json:"audio,omitempty"`
	} `json:"phonetics"`
	Origin   string `json:"origin"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
		} `json:"definitions"`
	} `json:"meanings"`
}

var DICTAPI = "https://api.dictionaryapi.dev/api/v2/entries/en/"
var Result []string

// prepares the url responsible to fetch the data from api
func prepareUrl(word string) {
	DICTAPI = DICTAPI + word
}

// makes http request to the prepared url and extracts the meaning, definition and example
func GetMeaning(word string) error {
	prepareUrl(word)
	response, err := http.Get(DICTAPI)
	if err != nil {
		Result = append(Result, "Network Error!", "Please check your internet connection!")
		return nil
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	Result = append(Result, strconv.Itoa(response.StatusCode))

	if response.StatusCode == 404 {
		Result = append(Result, "No Result :(", "Sorry mate, couldn't find the meaning for the word you specified!")
	} else if response.StatusCode != 200 {
		Result = append(Result, response.Status, "Some error occured! Please Try again later.")
	} else {
		var data apiSuccessData

		if err := json.Unmarshal([]byte(responseData), &data); err != nil {
			return err
		}
		Result = append(Result, data[0].Word, data[0].Meanings[0].Definitions[0].Definition, data[0].Phonetics[1].Audio)
	}
	return nil
}

func GetAudio() (io.Reader, error) {
	audio, err := http.Get(string(Result[len(Result)-1]))
	if err != nil {
		return nil, err
	}
	return audio.Body, nil

}
