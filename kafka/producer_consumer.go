package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "first_topic", 0)
	//conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	producer(conn)
	consumer(conn)
}

func producer(conn *kafka.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()
	conn.WriteMessages(kafka.Message{Value: []byte(message)})
}

func consumer(conn *kafka.Conn) {

	batch := conn.ReadBatch(1e2, 1e4)
	for {
		bytes := make([]byte, 1e4)
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
