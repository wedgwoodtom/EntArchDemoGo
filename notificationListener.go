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
	schedule(process, 100*time.Millisecond)
	n.listening = true
}

func (n *NotificationListener) Stop() {
	n.stopScheduler <- true
	n.listening = false
}

func (n *NotificationListener) processNotification() {
	// TODO: This is just a temp method as we are not actually listening to real notifications
	// Pretend we received a notification and so, put one on the Queue
	fmt.Println("Notification Received, writing Q message")

	n.NotificationQueue.PushItem()
}

//func init() {
//	fileName := "media-notification.txt"
//	dummyMediaNotification, err := ioutil.ReadFile(fileName)
//	if (err != nil) {
//		fmt.Println("File not found", fileName)
//	}
//	fmt.Println(string(dummyMediaNotification))
//}
