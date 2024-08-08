package notify

import (
	"fmt"
	"log"
	"time"
	
	d "github.com/vikySeeker/dicx/api"
	gn "git.sr.ht/~neftas/go-notify"
	//a "github.com/vikySeeker/dicx/audio"
	//gn "github.com/codegoalie/golibnotify"
)

// options value for the notify-send command to specify urgency and icon
var urgency string = "normal"
var success_icon string = "dicx"
var failure_icon string = "dicx-failed"
var Audio_source bool = true

func handlePlayAudio() {
	
}

func notifyUser(title string, message string, icon_path string) {

	//Attaching loggers to trace the notification flow...
	gn.SetDebugLogger(log.Println)
	gn.SetErrorLogger(log.Println)

	notifier, err:= gn.NewNotifier("Dicx go-notify")
	if err != nil {
		log.Println("Notifier: %v", err)
	}
	notification, err := notifier.NewNotification(title, message, icon_path)
	if err != nil {
		log.Println("Notification: %v", err)
	}

	err = notification.SetUrgency(gn.Normal)
	if err != nil {
		log.Println("Urgency: %v", err)
	}
	
	err = notification.SetTimeout(gn.DefaultTimeout)
	if err != nil {
		log.Println("Timeout: %v", err)
	}

	cb:= func() {
		fmt.Println("Button Clicked!")
	}
	if err = notification.AddAction("1", "Select", cb); err != nil {
		log.Println("Add Action: %v", err)
	}

	err = notification.Show()
	if err != nil {
		log.Println("Show: %v", err)
	}

	/*if err != nil {
		err = fmt.Errorf("Failed to send Notification: %w", err)
		log.Fatal(err)
	}*/

	time.Sleep(3*time.Second)

	notification.Close()
	/*if err != nil {
		err = fmt.Errorf("Failed to safely close the notification: %w", err)
                log.Fatal(err)
	}*/
}

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

	//cmd := exec.Command("notify-send", "-i", icon, "-u", urgency, "-A", "p=Read A Loud", word, meaning)
	_, err := d.GetAudio()
	if err != nil {
		//cmd = exec.Command("notify-send", "-i", icon, "-u", urgency, word, meaning)
		Audio_source = false
	}

	notifyUser(word, meaning, icon)
	/*
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
	*/
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
