package services

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {

    // Produce messages
    p := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "my-topic",
    })

    err := p.WriteMessages(context.Background(), 
        kafka.Message{
            Key: []byte("Key-A"),
            Value: []byte("Hello World!"),
        },
    )
    if err != nil {
        panic("could not write message " + err.Error())
    }

    // Consume messages
    c := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "my-topic",
    })

    for {
        m, err := c.ReadMessage(context.Background()) 
        if err != nil {
            break
        }
        fmt.Printf("message: %s\n", string(m.Value))
    }
}
