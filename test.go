package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// MessageQueue is a simple in-memory queue to store binary messages.
type MessageQueue struct {
	mu       sync.Mutex
	messages [][]byte
}

// Enqueue adds a binary message to the queue.
func (mq *MessageQueue) Enqueue(message []byte) {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	mq.messages = append(mq.messages, message)
}

// Dequeue removes and returns the oldest message in the queue.
func (mq *MessageQueue) Dequeue() ([]byte, bool) {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	if len(mq.messages) == 0 {
		return nil, false
	}
	msg := mq.messages[0]
	mq.messages = mq.messages[1:]
	return msg, true
}

var queue = &MessageQueue{}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		// Read command type (1 byte): 1 for PUBLISH, 2 for CONSUME
		command, err := reader.ReadByte()
		if err == io.EOF {
			log.Println("Client disconnected")
			return
		}
		if err != nil {
			log.Printf("Error reading command: %v", err)
			return
		}

		switch command {
		case 1: // PUBLISH
			// Read the length of the incoming message (4 bytes)
			var msgLen uint32
			err := binary.Read(reader, binary.BigEndian, &msgLen)
			if err != nil {
				log.Printf("Error reading message length: %v", err)
				return
			}

			// Read the message (binary)
			msg := make([]byte, msgLen)
			_, err = io.ReadFull(reader, msg)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				return
			}

			// Enqueue the binary message
			queue.Enqueue(msg)
			log.Println("Message published (binary)")

			// Acknowledge the client
			conn.Write([]byte("Message published\n"))

		case 2: // CONSUME
			// Dequeue the oldest message
			msg, ok := queue.Dequeue()
			if ok {
				// Send the message length first (4 bytes)
				msgLen := uint32(len(msg))
				binary.Write(conn, binary.BigEndian, msgLen)

				// Send the binary message
				conn.Write(msg)
				log.Println("Message consumed (binary)")
			} else {
				conn.Write([]byte("No messages available\n"))
			}

		default:
			log.Println("Unknown command")
			conn.Write([]byte("Unknown command\n"))
		}
	}
}

func main() {
	// Start listening for incoming TCP connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Binary Message Queue Server listening on :8080")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle connection in a new goroutine
		go handleConnection(conn)
	}
}
