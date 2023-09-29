package notify

import (
	"log"
	"os/exec"
)

var urgency string

func SendNotification(message [3]string) {
	summary := message[1]
	body := message[2]
	urgency = "normal"
	if message[0] != "200" {
		urgency = "critical"
	}
	cmd := exec.Command("notify-send", "-u", urgency, summary, body)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
