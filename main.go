package main

import (
	"gwclient/config"
	"gwclient/messages"
	"gwclient/requests"

	"log"
	"os"
)

func main() {
	config, err := config.ParseFromCommandLine()
	if err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Stress test client is started with the following config:\n")
	log.Printf("\tGateway URL       : %v\n", config.URL)
	log.Printf("\tMessage type      : %v\n", msgTypeToString(config.MessageType))
	log.Printf("\tMessage size      : %v\n", config.MessageSize)
	log.Printf("\tNumber of requests: %v\n", config.Requests)
	log.Printf("\tConcurrency number: %v\n", config.Concurrency)
	log.Printf("\tUse KeepAlive     : %v\n", config.KeepAlive)

	msgs := messages.Generate(&config)
	oks, fails := requests.Send(msgs, &config)

	log.Printf("Successful request: %v\n", oks)
	log.Printf("Failed request: %v\n", fails)

	if fails != 0 {
		os.Exit(1)
	}
	os.Exit(0)
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
