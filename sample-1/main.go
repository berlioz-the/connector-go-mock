package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/berlioz-the/connector-go"
)

func main() {
	log.Printf("---------- PEER MONITOR -----------------")

	berlioz.Service("app").Endpoint("client").MonitorAll(func(peers map[string]interface{}) {
		log.Printf("---------- PEER MONITOR -----------------")
		log.Printf("--- PEERS: %v\n", peers)
	})

	berlioz.Database("contacts").MonitorFirst(func(database interface{}) {
		log.Printf("---------- DATABASE MONITOR -----------------")
		log.Printf("--- DATABASE PEER: %v\n", database)
	})

	resp, body, err := berlioz.Service("app").Endpoint("client").Request().Get(context.Background(), "/")
	if err != nil {
		log.Printf("--- Response Error: %s\n", err)
	} else {
		log.Printf("--- Response Status Code: %s\n", resp.Status)
		log.Printf("--- Response Body: %s\n", body)
	}

	log.Printf("My Identity: %v\n", berlioz.Identity())

	params := &dynamodb.ScanInput{}
	result, err := berlioz.Database("contacts").DynamoDB().Scan(context.Background(), params)
	if err != nil {
		log.Printf("--- DynamoDB::Scan Error: %v\n", err)
	} else {
		log.Printf("--- DynamoDB::Scan Result: %v\n", result)
	}

	time.Sleep(5 * time.Second)
}
