package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
	"time"
)

func schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func sqsNameToUrl(svc *sqs.SQS) {

}


func writeItem() {
	// Initialize a session in us-west-2 that the SDK will use to load credentials
	// from the shared config file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		fmt.Println("Error getting session:")
		fmt.Println(err)
		os.Exit(1)
	}

	// Put item on Q
	svc := sqs.New(sess)

	//svc.GetQueueUrl()
	name := "media-notifications"
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		fmt.Println("Unable to find queue: ", name)
		return
	}

	// URL to our queue - can use APIs to lookup the URL by name in order to make it region agnostic
	qURL := resultURL.QueueUrl

	msgWritten, err := svc.SendMessage(&sqs.SendMessageInput{
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
		QueueUrl:    qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Wrote 1 Message", *msgWritten.MessageId)
}


func readItem() []*sqs.Message {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		fmt.Println("Error getting session:")
		fmt.Println(err)
		os.Exit(1)
	}

	// Put item on Q
	svc := sqs.New(sess)

	//svc.GetQueueUrl()
	name := "media-notifications"
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		fmt.Println("Unable to find queue: ", name)
		return nil
	}

	// URL to our queue - can use APIs to lookup the URL by name in order to make it region agnostic
	qURL := resultURL.QueueUrl
	// Read item from Q

	msgRead, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            qURL,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(20),  // 20 seconds
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
		fmt.Println(msgRead.Messages)
	}

	return msgRead.Messages
}


func deleteItem(receiptHandle *string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		fmt.Println("Error getting session:")
		fmt.Println(err)
		os.Exit(1)
	}

	// Put item on Q
	svc := sqs.New(sess)

	//svc.GetQueueUrl()
	name := "media-notifications"
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		fmt.Println("Unable to find queue: ", name)
		return
	}

	// URL to our queue - can use APIs to lookup the URL by name in order to make it region agnostic
	qURL := resultURL.QueueUrl

	// Delete the item
	resultDelete, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      qURL,
		ReceiptHandle: receiptHandle,
	})

	// msgRead.Messages[0].ReceiptHandle

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}

	fmt.Println("Message Deleted", resultDelete)
}

func main() {

	writeItem()
	message := readItem()[0]
	deleteItem(message.ReceiptHandle)


	// Some Stuff

	//ping := func() { fmt.Println("#") }

	stop := schedule(doSomeWork, 2*time.Second)
	time.Sleep(10 * time.Second)
	stop <- true
	time.Sleep(10 * time.Second)

	fmt.Println("Done")
}


func doSomeWork() {
	fmt.Println("Doing some work")
	writeItem()
	message := readItem()[0]
	deleteItem(message.ReceiptHandle)
}




/*

ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
        for range ticker.C {
            fmt.Println("Tick")
        }
    }()
time.Sleep(1600 * time.Millisecond)
ticker.Stop()
 */