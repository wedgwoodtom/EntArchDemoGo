package main

import (
	"fmt"
	"time"
)

type NotificationListener struct {
	NotificationQueue SqsMessageQueue
	listening         bool
}

func (n *NotificationListener) Start() {
	// For this demo, just listen periodically since this is faking it
	n.listening = true

	for {
		// Pretend we received a notification and so, put one on the Queue
		n.processNotification()

		if !n.listening {
			return
		}
		time.After(1 * time.Second)
	}
}

func (n *NotificationListener) Stop() {
	n.listening = false
}

func (n *NotificationListener) processNotification() {
	// Pretend we received a notification and so, put one on the Queue
	fmt.Println("Notification Received, writing Q message")
	n.NotificationQueue.PushItem()
}
