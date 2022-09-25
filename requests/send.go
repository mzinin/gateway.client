package requests

import (
	"gwclient/config"

	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	jsonType = "application/json"
)

func Send(messages [][][]byte, cfg *config.Config) (uint32, uint32) {
	log.Printf("Sending requests\n")
	start := time.Now()

	oks, fails := sendRequests(messages, cfg.URL, cfg.KeepAlive)

	log.Printf("Done in %v\n", time.Since(start))
	return oks, fails
}

func sendRequests(messages [][][]byte, url string, keepAlive bool) (uint32, uint32) {
	oks := uint32(0)
	fails := uint32(0)

	wg := sync.WaitGroup{}
	for _, bucket := range messages {
		wg.Add(1)
		go func(bucket [][]byte) {
			defer wg.Done()

			transport := http.Transport{DisableKeepAlives: !keepAlive}
			client := http.Client{Transport: &transport}

			for _, data := range bucket {
				resp, err := client.Post(url, jsonType, bytes.NewBuffer(data))
				if err != nil {
					log.Printf("Error sending POST request: %v\n", err)
					atomic.AddUint32(&fails, uint32(1))
				} else {
					defer resp.Body.Close()
					io.Copy(ioutil.Discard, resp.Body)
					if resp.StatusCode/100 == 2 {
						atomic.AddUint32(&oks, uint32(1))
					} else {
						atomic.AddUint32(&fails, uint32(1))
					}
				}
			}
		}(bucket)
	}
	wg.Wait()

	return oks, fails
}
