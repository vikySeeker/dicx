package api

import (
	"fmt"
	"encoding/json"
	"errors"
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

// makes http request to the prepared url and extracts the meaning, definition and example
func GetMeaning(word string) error {
	DICTAPI = DICTAPI + word
	response, err := http.Get(DICTAPI)
	if err != nil {
		Result = append(Result, "69", "Network Error!", "Please check your internet connection!\nStatus: 69")
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
		Result = append(Result, response.Status, "Some error occurred! Please Try again later.")
	} else {
		var data apiSuccessData

		if err := json.Unmarshal([]byte(responseData), &data); err != nil {
			return err
		}

		url := audioURL(&data)
		if url != "" {
			Result = append(Result, data[0].Word, data[0].Meanings[0].Definitions[0].Definition, url)
		} else {
			Result = append(Result, data[0].Word, data[0].Meanings[0].Definitions[0].Definition, "nil")
		}
	}
	return nil
}

func audioURL(meaning *apiSuccessData) string {
	for _, data := range (*meaning)[0].Phonetics {
		if data.Audio != "" {
			return data.Audio
		}
	}
	return ""
}

func GetAudio() (io.Reader, error) {
	url := Result[len(Result)-1]
	if url != "nil" {
		audio, err := http.Get(string(Result[len(Result)-1]))
		if err != nil {
			return nil, err
		}
		return audio.Body, nil
	}
	return nil, errors.New("no audio source found")
}
