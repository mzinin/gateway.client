package main

import (
	"gwclient/config"

	"fmt"
	"log"
	"os"
)

func main() {
	config, err := config.ParseFromCommandLine()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Stress test client is started with the following config:\n")
	log.Printf("\tGateway URL       : %v\n", config.URL)
	log.Printf("\tMessage type      : %v\n", msgTypeToString(config.MessageType))
	log.Printf("\tMessage size      : %v\n", config.MessageSize)
	log.Printf("\tNumber of requests: %v\n", config.Requests)
	log.Printf("\tConcurrency number: %v\n", config.Concurrency)
	log.Printf("\tUse KeepAlive     : %v\n", config.KeepAlive)
}

func msgTypeToString(msgType config.MessageType) string {
	switch msgType {
	case config.Json:
		return "JSON"
	case config.FlatBuffers:
		return "FlatBuffers"
	}
	return ""
}
