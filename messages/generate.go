package messages

import (
	"gwclient/config"

	"log"
	"sync"
	"time"
)

func Generate(cfg *config.Config) [][][]byte {
	log.Printf("Generating %v messages\n", cfg.Requests)
	start := time.Now()

	messages := generateMessages(cfg.Concurrency, cfg.Requests)

	var result [][][]byte
	if cfg.MessageType == config.Json {
		result = convertToJson(messages, cfg.MessageSize)
	} else {
		result = convertToFbf(messages, cfg.MessageSize)
	}

	log.Printf("Done in %v\n", time.Since(start))
	return result
}

func generateMessages(buckets uint16, total uint32) [][]*message {
	bucketsSizes := calculateBucketsSizes(buckets, total)

	result := make([][]*message, buckets)

	start := uint32(0)
	for i, size := range bucketsSizes {
		result[i] = generateBucket(size, start)
		start += size
	}

	return result
}

func calculateBucketsSizes(buckets uint16, total uint32) []uint32 {
	result := make([]uint32, buckets)

	minSize := total / uint32(buckets)
	for i, _ := range result {
		result[i] = minSize
		total -= minSize
	}
	for i := uint32(0); i < total; i++ {
		result[i] += 1
	}

	return result
}

func generateBucket(size uint32, start uint32) []*message {
	result := make([]*message, size)

	for i, _ := range result {
		result[i] = createMessage()
		result[i].Int0 = int64(start % 1000)
		result[i].Int1 = int64((start / 1000) % 1000)
		result[i].Int2 = int64((start / 1000000) % 1000)
		result[i].Int3 = int64((start / 1000000000) % 1000)
		start += 1
	}

	return result
}

func convertToJson(messages [][]*message, targetSize uint16) [][][]byte {
	result := make([][][]byte, len(messages))

	wg := sync.WaitGroup{}
	for i, bucket := range messages {
		wg.Add(1)
		go func(i int, bucket []*message) {
			defer wg.Done()
			result[i] = make([][]byte, len(bucket))
			for j, message := range bucket {
				currentSize := message.jsonSize()
				if currentSize < int(targetSize) {
					message.increaseSize(uint(targetSize) - uint(currentSize))
				}
				result[i][j] = message.toJson()
			}
		}(i, bucket)
	}
	wg.Wait()

	return result
}

func convertToFbf(messages [][]*message, targetSize uint16) [][][]byte {
	result := make([][][]byte, len(messages))

	wg := sync.WaitGroup{}
	for i, bucket := range messages {
		wg.Add(1)
		go func(i int, bucket []*message) {
			defer wg.Done()
			result[i] = make([][]byte, len(bucket))
			for j, message := range bucket {
				currentSize := message.fbfSize()
				if currentSize < int(targetSize) {
					message.increaseSize(uint(targetSize) - uint(currentSize))
				}
				result[i][j] = message.toFbf()
			}
		}(i, bucket)
	}
	wg.Wait()

	return result
}
