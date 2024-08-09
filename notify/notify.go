package notify

import (
	"fmt"
	"log"
	"time"
	"io"
	
	d "github.com/vikySeeker/dicx/api"
	gn "git.sr.ht/~neftas/go-notify"
	a "github.com/vikySeeker/dicx/audio"
)

// options value for the libnotify notification to specify urgency and icon.
var is_critical bool = false
var success_icon string = "/home/seeker/projects/personal/dicx/icons/dicx.png"
var failure_icon string = "/home/seeker/projects/personal/dicx/icons/dicx-failed.png"
var audio_source bool = true
var notified bool = false

/* 
function that is responsible for generating notification using libnotify.
*/
func notifyUser(title string, message string, icon_path string, audiodata *io.Reader) {

	//Attaching loggers to trace the notification flow...
	//gn.SetDebugLogger(log.Println)
	//gn.SetErrorLogger(log.Println)

	notifier, err:= gn.NewNotifier("Dicx!")
	if err != nil {
		log.Printf("Notifier: %v\n", err)
	}
	notification, err := notifier.NewNotification(title, message, icon_path)
	if err != nil {
		log.Printf("Notification: %v\n", err)
	}

	if is_critical {
		err = notification.SetUrgency(gn.Critical)
	} else { 
		err = notification.SetUrgency(gn.Normal)
	}

	if err != nil {
		log.Printf("Urgency: %v\n", err)
	}
	
	playAudioHandler := func() {
		err := a.PronounceWord(audiodata)
		if err != nil {
			log.Printf("Playing Audio: %v\n", err)
		}
		notified = true
	}

	closeNotificationHandler := func() {
		notified = true
		err := notification.Close()
		if err != nil {
			err = fmt.Errorf("Failed to safely close the notification: %w", err)
	                log.Fatal(err)
		}
		fmt.Println("Notification Close Successfully!!\n")
	}

	if audio_source && audiodata != nil {
		err = notification.SetTimeout(gn.InfinityTimeout)
		if err != nil {
			log.Printf("Timeout: %v\n", err)
		}
		if err = notification.AddAction("1", "Audio", playAudioHandler); err != nil {
			log.Printf("Add Action: %v\n", err)
		}
		if err = notification.AddAction("2", "Close", closeNotificationHandler); err != nil {
			log.Printf("Add Action: %v\n", err)
		}
	} else {
		err = notification.SetTimeout(gn.DefaultTimeout)
		if err != nil {
			log.Printf("Timeout: %v\n", err)
		}
	}
	
	err = notification.Show()
	if err != nil {
		log.Println("Show: %v", err)
	}
	
	timer := 0
	for audio_source && !notified {
		if timer == 10 {
			notification.Close()
			break
		}
		time.Sleep(1 * time.Second)
		timer += 1
		continue
	}
}

/*
function that is responsible for sending notification
*/
func SendNotification(message []string) {
	word := message[1]
	meaning := message[2]
	icon := success_icon
	if message[0] != "200" {
		is_critical = true
		icon = failure_icon
	}

	audiodata, err := d.GetAudio()
	if err != nil {
		audio_source = false
		notifyUser(word, meaning, icon, nil)

	} else {
		notifyUser(word, meaning, icon, &audiodata)
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
