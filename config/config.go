package config

type MessageType int

const (
	Json MessageType = iota
	FlatBuffers
)

type Config struct {
	MessageType      MessageType
	MessageSize      uint16
	NumberOfMessages uint32
	Concurrency      uint16
	KeepAlive        bool
	URL              string
}

func makeDefaultConfig() Config {
	return Config{
		MessageType:      Json,
		MessageSize:      128,
		Concurrency:      1,
		NumberOfMessages: 10000,
		KeepAlive:        false,
		URL:              "",
	}
}
