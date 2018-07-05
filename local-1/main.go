package main

import (
	"fmt"
	"log"
	"net/http"

	"connector-go.git"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	log.Printf("---------- PEER MONITOR -----------------")

	// time.Sleep(1 * time.Second)

	// berlioz.TestZipkin()

	// time.Sleep(2 * time.Second)

	http.HandleFunc("/", berlioz.WrapFunc(sayhelloName)) // set router
	err := http.ListenAndServe(":4000", nil)             // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")

	fmt.Fprintf(w, "From APP: ")
	_, body, err := berlioz.Request("service", "app", "client").Get(r.Context(), "/")
	if err != nil {
		fmt.Fprintf(w, "Error: ")
		fmt.Fprintf(w, "%#v", err)
	} else {
		fmt.Fprintf(w, string(body[:]))
	}
	fmt.Fprintf(w, "\n")

	params := &dynamodb.ScanInput{}
	result, err := berlioz.Database("contacts").DynamoDB().Scan(r.Context(), params)
	if err != nil {
		fmt.Fprintf(w, "--- DynamoDB::Scan Error: %v\n", err)
	} else {
		fmt.Fprintf(w, "--- DynamoDB::Scan Result: %v\n", result)
	}
	fmt.Fprintf(w, "\n")
}
