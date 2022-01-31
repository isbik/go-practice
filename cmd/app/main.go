package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {

	conn, err := kafka.Dial("tcp", "localhost:9092")

	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	partitions, err := conn.ReadPartitions()

	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	for _, k := range m {
		fmt.Println(k)
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset,
			string(m.Key), string(m.Value))
	}
	
	if err := r.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}
}
