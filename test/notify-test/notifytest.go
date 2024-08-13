package main

import (
	"fmt"
	"log"
	"time"

	gn "github.com/codegoalie/golibnotify"
)

func main() {
	notifier := gn.NewSimpleNotifier("Dicx Libnotify Test")
	err := notifier.Show("Testing", "Testing libnotify in golang for dicx", "/home/seeker/projects/personal/dicx/icons/dicx.png")
	if err != nil {
		err = fmt.Errorf("Failed to send a notification: %w", err)
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	err = notifier.Close()
	if err != nil {
		err = fmt.Errorf("Failed to safely close the notification: %w", err)
		log.Fatal(err)
		
	}

}
