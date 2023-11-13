package notify

import (
	"fmt"
	"os/exec"
)

// options value for the notify-send command to specify urgency and icon
var urgency string = "normal"
var success_icon string = "dicx"
var failure_icon string = "dicx-failed"

/*
function that is responsible for sending notification
*/
func SendNotification(message [3]string) error {
	word := message[1]
	meaning := message[2]
	icon := success_icon
	if message[0] != "200" {
		urgency = "critical"
		icon = failure_icon
	}
	cmd := exec.Command("notify-send", "-i", icon, "-u", urgency, word, meaning)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

/*
function that is responsible for printing the output to terminal
*/
func PrintOutput(message [3]string) error {
	word := message[1]
	meaning := message[2]
	_, err := fmt.Println("Selected Word is: ", word, "\nMeaning:", meaning)
	if err != nil {
		return err
	}
	return nil
}
