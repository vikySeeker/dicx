package main

import (
	"fmt"
	"log"
	"time"

	"git.sr.ht/~neftas/go-notify"
)

func main() {
	// Optionally attach loggers to see what's going on
	notify.SetDebugLogger(log.Println)
	notify.SetErrorLogger(log.Println)

	// Get an instance of a Notifier
	notifier, err := notify.NewNotifier("Example app")
	if err != nil {
		log.Fatalf("NewNotifier: %v", err)
	}

	// Create and show 2 notifications
	criticalUrgencyNotification, err := notifier.NewNotification(
		"First notification",
		"This has critical urgency",
		"a/path/to/an/icon.png",
	)
	if err != nil {
		log.Fatalf("NewNotification: %v", err)
	}
	if err := criticalUrgencyNotification.SetUrgency(notify.Critical); err != nil {
		log.Fatalf("set urgency on criticalUrgencyNotification: %v", err)
	}
	if err := criticalUrgencyNotification.SetTimeout(notify.InfinityTimeout); err != nil {
		log.Fatalf("set timeout on criticalUrgencyNotification: %v", err)
	}
	cb := func() {
		fmt.Println("Do something from withing the callback")
	}
	if err := criticalUrgencyNotification.AddAction("1", "Accept", cb); err != nil {
		log.Fatalf("add action: %v", err)
	}
	if err := criticalUrgencyNotification.AddAction("2", "Reject", cb); err != nil {
		log.Fatalf("add action: %v", err)
	}
	if err := criticalUrgencyNotification.Show(); err != nil {
		log.Fatalf("show criticalUrgencyNotification: %v", err)
	}

	// Urgency.Normal is the default
	normalUrgencyNotification, err := notifier.NewNotification(
		"Another notification",
		"This one with normal urgency",
		"a/path/to/an/icon.png",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := normalUrgencyNotification.SetTimeout(2000); err != nil {
		log.Fatalf("set timeout on normalUrgencyNotification: %v", err)
	}
	if err := normalUrgencyNotification.Show(); err != nil {
		log.Fatalf("show normalUrgencyNotification: %v", err)
	}

	time.Sleep(3 * time.Second)

	// Update an existing criticalUrgencyNotification (or send a new one if one hasn't
	// been sent)
	if err := criticalUrgencyNotification.Update(
		"Updated first notification",
		"It should now have low urgency",
		"another/path/to/icon.png",
	); err != nil {
		log.Fatalf("update criticalUrgencyNotification: %v", err)
	}
	// Also update urgency to Low
	if err := criticalUrgencyNotification.SetUrgency(notify.Low); err != nil {
		log.Fatalf("set urgency to low on criticalUrgencyNotification: %v", err)
	}
	if err := criticalUrgencyNotification.SetTimeout(10000); err != nil {
		log.Fatalf("set timeout to 10 secs on criticalUrgencyNotification: %v", err)
	}
	if err := criticalUrgencyNotification.Show(); err != nil {
		log.Fatalf("show criticalUrgencyNotification: %v", err)
	}

	time.Sleep(3 * time.Second)

	// Remove an existing criticalUrgencyNotification
	if err := criticalUrgencyNotification.Close(); err != nil {
		err = fmt.Errorf("failed to close criticalUrgencyNotification: %w", err)
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)
}
