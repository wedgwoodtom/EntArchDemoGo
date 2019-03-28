package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	// Dyn

	// Start Notification Listener and Compute Service
	notificationQueue := NewQueue("us-west-2", "media-notifications")

	notificationListener := NotificationListener{NotificationQueue: notificationQueue}
	notificationListener.Start()

	computeService := ComputeService{notificationQueue: notificationQueue}
	computeService.Start()

	// Start web service
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	// Sleep/Run for 5 seconds
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
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
