package main


import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"log"
)

func initPubSub() {
	ctx := context.Background()

	projectID := "gotest2-jl1243232"

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Couldn't create client: %v", err)
	}

	topicID := "test-topic"

	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Failure when creating topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}