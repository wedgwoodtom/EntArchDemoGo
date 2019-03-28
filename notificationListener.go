package main

import (
	"fmt"
	"time"
)

type NotificationListener struct {
	NotificationQueue SqsMessageQueue
	stopScheduler     chan bool
	listening         bool
}

func (n *NotificationListener) Start() {
	// For this demo, just listen periodically since this is faking it
	process := func() {
		n.processNotification()
	}
	Schedule(process, 1*time.Second)
	n.listening = true
}

func (n *NotificationListener) Stop() {
	n.stopScheduler <- true
	n.listening = false
}

func (n *NotificationListener) processNotification() {
	// Pretend we received a notification and so, put one on the Queue
	fmt.Println("Notification Received, writing Q message")
	n.NotificationQueue.PushItem()
}
