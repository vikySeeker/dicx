package notify

import (
	d "dicx/api"
	a "dicx/audio"
	"fmt"
	"log"
	"os/exec"
)

// options value for the notify-send command to specify urgency and icon
var urgency string = "normal"
var success_icon string = "dicx"
var failure_icon string = "dicx-failed"

/*
function that is responsible for sending notification
*/
func SendNotification(message []string) {
	word := message[1]
	meaning := message[2]
	icon := success_icon
	if message[0] != "200" {
		urgency = "critical"
		icon = failure_icon
	}
	cmd := exec.Command("notify-send", "-i", icon, "-u", urgency, "-A", "p=Spell", word, meaning)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	if string(output[0]) == "p" {
		data, err := d.GetAudio()
		if err != nil {
			log.Fatal(err)
		}
		err = a.PronounceWord(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*
function that is responsible for printing the output to terminal
*/
func PrintOutput(message []string) {
	word := message[1]
	meaning := message[2]
	_, err := fmt.Println("Selected Word is: ", word, "\nMeaning:", meaning)
	if err != nil {
		log.Fatal(err)
	}
}
