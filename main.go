package main

import (
	"gwclient/config"

	"fmt"
	"os"
)

func main() {
	config, err := config.ParseFromCommandLine()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("config: %+v\n", config)
}
