package main

import (
	"log"
	"time"

	"connector-go.git"
)

func main() {
	log.Printf("---------- PEER MONITOR -----------------")

	time.Sleep(1 * time.Second)

	berlioz.TestZipkin()

	time.Sleep(2 * time.Second)
}
