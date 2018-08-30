package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"connector-go.git"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/***************************************************************/

type serviceInfo struct {
	id        string
	serviceID string
	endpoint  string
	handler   func(map[string]interface{})
	monitor   berlioz.SubscribeInfo
}

var monitoredAgents = make(map[string]serviceInfo)
var newAgents = make(map[string]bool)
var trackedPeers = make(map[string]map[string]interface{})

func forever() {
	for {
		// fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}

func processPeers() {
	fmt.Printf("***** TRACKED PEERS: %#v\n", trackedPeers)

}

func monitorAgent(consumed berlioz.ConsumesModel) {
	fmt.Printf("***** AGENT TO MONITOR: %s\n", consumed.ID)
	id := consumed.ID + "-" + consumed.Endpoint
	info := serviceInfo{id: id, serviceID: consumed.ID, endpoint: consumed.Endpoint}
	info.handler = func(peers map[string]interface{}) {
		fmt.Printf("***** PEERS CHANGED FOR: %s\n", id)
		trackedPeers[id] = peers
		processPeers()
	}

	monitoredAgents[id] = info
	newAgents[id] = true

	info.monitor = berlioz.Sector(consumed.Sector).Service(consumed.Name).Endpoint(consumed.Endpoint).MonitorAll(info.handler)
}

func stopMonitoring(id string, serviceInfo serviceInfo) {
	fmt.Printf("***** AGENT TO STOP MONITORING: %s\n", id)
	serviceInfo.monitor.Stop()
	delete(trackedPeers, id)
	delete(monitoredAgents, id)
	processPeers()
}

func applyAgentChanges() {
	// fmt.Printf("***** applyAgentChanges: %#v\n", monitoredAgents)

	for id, serviceInfo := range monitoredAgents {
		if _, ok := newAgents[id]; !ok {
			stopMonitoring(id, serviceInfo)
		}
	}
}

func onConsumesChanged(consumes []berlioz.ConsumesModel) {
	// fmt.Printf("***** UPDATED MONITOR CONSUMES: %#v\n", consumes)
	newAgents = make(map[string]bool)
	for _, consumed := range consumes {
		if consumed.Kind == "service" && consumed.Name == "berlioz_agent" && consumed.Endpoint == "ws" {
			monitorAgent(consumed)
		}
	}
	applyAgentChanges()
}

func testPrometheus() {
	berlioz.Consumes().MonitorAll(onConsumesChanged)

	forever()
}

/***************************************************************/

func main() {
	testPrometheus()
	return

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
}
