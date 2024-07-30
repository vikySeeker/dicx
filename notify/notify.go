package notify

import (
	d "github.com/vikySeeker/dicx/api"
	a "github.com/vikySeeker/dicx/audio"
	"fmt"
	"log"
	"os/exec"
)

// options value for the notify-send command to specify urgency and icon
var urgency string = "normal"
var success_icon string = "dicx"
var failure_icon string = "dicx-failed"
var Audio_source bool = true

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

	cmd := exec.Command("notify-send", "-i", icon, "-u", urgency, "-A", "p=Read A Loud", word, meaning)
	data, err := d.GetAudio()
	if err != nil {
		cmd = exec.Command("notify-send", "-i", icon, "-u", urgency, word, meaning)
		Audio_source = false
	}
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	if len(output) > 0 && Audio_source {
		if string(output[0]) == "p" {
			err = a.PronounceWord(data)
			if err != nil {
				log.Fatal(err)
			}
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
