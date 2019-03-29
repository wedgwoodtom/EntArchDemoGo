package main

import (
	"fmt"
	"testing"
)

func TestPut(t *testing.T) {
	media := MakeRandomMedia()
	dynamo := NewDynamoDb("us-west-2")
	dynamo.PutItem(media, "MediaNotification")
	fmt.Println("item put")
}
