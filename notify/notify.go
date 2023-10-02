package notify

import (
	"fmt"
	"os/exec"
)

var urgency string

/*
function that is responsible for sending notification
*/
func SendNotification(message [3]string) error {
	word := message[1]
	meaning := message[2]
	urgency = "normal"
	if message[0] != "200" {
		urgency = "critical"
	}
	cmd := exec.Command("notify-send", "-u", urgency, word, meaning)
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
