package main

import (
	"context"
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
	_, body, err := berlioz.Service("app").Request().Get(context.Background(), "/")
	if err != nil {
		fmt.Printf("***** REQUEST Error: ")
		fmt.Printf("%#v", err)
	} else {
		fmt.Printf("***** REQUEST RESULT: ")
		fmt.Printf(string(body[:]))
	}
	fmt.Printf("\n")

	berlioz.MyEndpoint("default").Monitor(func(ep berlioz.EndpointModel) {
		fmt.Printf("**** MONITOR DEFAULT EP: %#v. Present: %t\n", ep, ep.IsPresent())
	})

	berlioz.Service("app").MonitorAll(func(peers map[string]interface{}) {
		fmt.Printf("***** UPDATED APP PEERS: %#v\n", peers)
		fmt.Printf("***** UPDATED APP PEERS MANUAL GET: %#v\n", berlioz.Service("app").All())
	})

	berlioz.Service("app").MonitorFirst(func(peer interface{}) {
		fmt.Printf("***** UPDATED APP FIRST PEER: %#v\n", peer)
	})

	http.HandleFunc("/", berlioz.WrapFunc(sayhelloName)) // set router
	err = http.ListenAndServe(":4000", nil)              // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")

	fmt.Fprintf(w, "From APP: ")
	_, body, err := berlioz.Service("app").Request().Get(r.Context(), "/")
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

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")

	ep := berlioz.MyEndpoint("default").Get()
	fmt.Fprintf(w, "GOSSIP EP: %v\n", ep)

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "My Identity: %v\n", berlioz.Identity())

}
