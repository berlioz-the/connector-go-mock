package main

import (
	"log"
	"time"

	"connector-go.git"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	log.Printf("---------- PEER MONITOR -----------------")

	berlioz.Peers("service", "app", "client").Monitor(func(peers berlioz.PeerAccessor) {
		log.Printf("---------- PEER MONITOR -----------------")
		log.Printf("--- PEERS: %v\n", peers.All())
		if val, ok := peers.Get("1"); ok {
			log.Printf("--- INDEXED PEER: %v\n", val)
		}
		if val, ok := peers.Random(); ok {
			log.Printf("--- RANDOM PEER: %v\n", val)
		}
	})

	berlioz.Database("contacts").Monitor(func(database berlioz.NativeResourceAccessor) {
		log.Printf("---------- DATABASE MONITOR -----------------")

		params := &dynamodb.ScanInput{}
		result, err := database.DynamoDB().Scan(params)
		if err != nil {
			log.Printf("--- DynamoDB::Scan Error: %v\n", err)
		} else {
			log.Printf("--- DynamoDB::Scan Result: %v\n", result)
		}
	})

	resp, body, err := berlioz.Request("service", "app", "client").Get("/")
	if err != nil {
		log.Printf("--- Response Error: %s\n", err)
	} else {
		log.Printf("--- Response Status Code: %s\n", resp.Status)
		log.Printf("--- Response Body: %s\n", body)
	}

	berlioz.TestZipkin()

	time.Sleep(5 * time.Second)
}
