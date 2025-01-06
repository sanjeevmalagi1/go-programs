package main

import (
	"fmt"
	"time"
)

type PubSub struct {
	subscribers map[string][]chan string // Topic to subscribers map
}

// NewPubSub initializes the PubSub system.
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

// Subscribe adds a subscriber to a topic and returns a channel to receive messages.
func (ps *PubSub) Subscribe(topic string) <-chan string {
	ch := make(chan string)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

// Publish sends a message to all subscribers of a topic.
func (ps *PubSub) Publish(topic, message string) {
	if chans, found := ps.subscribers[topic]; found {
		for _, ch := range chans {
			ch <- message
		}
	}
}

func main() {
	ps := NewPubSub()

	// Subscriber 1 subscribes to "news"
	sub1 := ps.Subscribe("news")
	go func() {
		for msg := range sub1 {
			fmt.Println("Subscriber 1 received:", msg)
		}
	}()

	// Subscriber 2 subscribes to "news"
	sub2 := ps.Subscribe("news")
	go func() {
		for msg := range sub2 {
			fmt.Println("Subscriber 2 received:", msg)
		}
	}()

	// Publish messages
	ps.Publish("news", "Breaking News: Go is awesome!")
	time.Sleep(1 * time.Second)

	ps.Publish("news", "Another news update!")
	time.Sleep(1 * time.Second)
}
