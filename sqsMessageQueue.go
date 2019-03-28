package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

type SqsMessageQueue struct {
	session  session.Session
	service  *sqs.SQS
	queueUrl string
}

func NewQueue(region string, queueName string) SqsMessageQueue {

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		fmt.Println("Error getting session:")
		fmt.Println(err)
		os.Exit(1)
	}

	// Put item on Q
	service := sqs.New(session)

	//svc.GetQueueUrl()
	name := queueName
	resultURL, err := service.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		fmt.Println("Unable to find queue: ", name)
		os.Exit(1)
	}

	queueUrl := resultURL.QueueUrl

	return SqsMessageQueue{session: *session, service: service, queueUrl: *queueUrl}

}

func (queue *SqsMessageQueue) PullItems() []*sqs.Message {

	msgRead, err := queue.service.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queue.queueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
		return nil
	}

	if len(msgRead.Messages) == 0 {
		fmt.Println("Received no messages")
		return nil
	} else {
		fmt.Println("Received messages", len(msgRead.Messages))
		//fmt.Println(msgRead.Messages)
	}

	return msgRead.Messages
}

func (queue *SqsMessageQueue) PushItem() {
	msgWritten, err := queue.service.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("Information about current NY Times fiction bestseller."),
		QueueUrl:    &queue.queueUrl,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Wrote 1 Message", *msgWritten.MessageId)
}

func (queue *SqsMessageQueue) deleteItem(receiptHandle *string) {
	// Delete the item
	resultDelete, err := queue.service.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &queue.queueUrl,
		ReceiptHandle: receiptHandle,
	})

	// msgRead.Messages[0].ReceiptHandle

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}

	resultDelete.GoString()
	fmt.Println("Message Deleted for receipt", receiptHandle)
}
