package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type DynamoDb struct {
	service *dynamodb.DynamoDB
}

func NewDynamoDb(region string) DynamoDb {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	service := dynamodb.New(sess)

	return DynamoDb{service: service}
}

func (d *DynamoDb) PutItem(object interface{}, table string) {
	av, err := dynamodbattribute.MarshalMap(object)
	if err != nil {
		fmt.Println("Got error marshalling object:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Println("av ", av)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = d.service.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
