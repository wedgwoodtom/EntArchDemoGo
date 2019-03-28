package main

import (
	"fmt"
	"time"
)

type ComputeService struct {
	notificationQueue SqsMessageQueue
}

func New(notificationQueue SqsMessageQueue) ComputeService {
	return ComputeService{notificationQueue: notificationQueue}
}

func (cs *ComputeService) Start() {
	// For this demo, just listen periodically since this is faking it

	process := func() {
		cs.processNotificationMessages()
	}
	Schedule(process, 1*time.Second)
}

func (cs *ComputeService) processNotificationMessages() {
	messages := cs.notificationQueue.PullItems()
	if messages == nil {
		return
	}

	for _, message := range messages {
		fmt.Println("Computing for Message: ", message.MessageId)

		// Store to Dynamo

		cs.notificationQueue.deleteItem(message.ReceiptHandle)
	}

}
