package config

import (
	"flag"
	"fmt"
	"net/url"
)

var (
	msgType     string
	msgSize     uint
	requests    uint
	concurrency uint
	keepAlive   bool
	gatewayURL  string
)

func ParseFromCommandLine() (Config, error) {
	result := makeDefaultConfig()

	parseCommandLineKeys()

	if msgType == "json" {
		result.MessageType = Json
	} else if msgType == "fbf" {
		result.MessageType = FlatBuffers
	} else {
		return result, fmt.Errorf("Unsupported message type: %s", msgType)
	}

	if msgSize > 65535 {
		return result, fmt.Errorf("Message size is too big (max 65535): %d", msgSize)
	} else {
		result.MessageSize = uint16(msgSize)
	}

	if requests > 4294967295 {
		return result, fmt.Errorf("Number of requests in too big (max 4294967295): %d", requests)
	} else {
		result.Requests = uint32(requests)
	}

	if concurrency > 65535 {
		return result, fmt.Errorf("Concurrency number in too big (max 65535): %d", concurrency)
	} else {
		result.Concurrency = uint16(concurrency)
	}

	url, err := url.Parse(gatewayURL)
	if err != nil {
		return result, fmt.Errorf("Wrong gateway URL %v: %v", gatewayURL, err)
	} else if url.Scheme != "http" {
		return result, fmt.Errorf("Wrong gateway URL %v: unsupported sheme", gatewayURL)
	} else if len(url.Host) == 0 {
		return result, fmt.Errorf("Wrong gateway URL %v: empty host", gatewayURL)
	} else {
		result.URL = gatewayURL
	}

	result.KeepAlive = keepAlive

	return result, nil
}

func parseCommandLineKeys() {
	// message type
	flag.StringVar(&msgType, "t", "json", "Message type: json | fbf")

	// message size
	flag.UintVar(&msgSize, "s", 512, "Message size")

	// number of messages
	flag.UintVar(&requests, "n", 10000, "Number of requests to perform")

	// concurrency
	flag.UintVar(&concurrency, "c", 1, "Number of multiple requests to make at a time")

	// keep alive flag
	flag.BoolVar(&keepAlive, "k", false, "Use HTTP KeepAlive feature")

	// gateway url
	flag.StringVar(&gatewayURL, "u", "", "Gateway URL")

	flag.Parse()
}
