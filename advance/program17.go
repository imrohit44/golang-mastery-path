package main

import (
	"fmt"
	"sync"
	"time"
)

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan string
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan string),
	}
}

func (eb *EventBus) Subscribe(topic string) <-chan string {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	ch := make(chan string, 1)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

func (eb *EventBus) Publish(topic string, data string) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	
	if chans, found := eb.subscribers[topic]; found {
		for _, ch := range chans {
			// Non-blocking send
			select {
			case ch <- data:
			default:
			}
		}
	}
}

func main() {
	bus := NewEventBus()

	userEvents := bus.Subscribe("user_created")

	go func() {
		for msg := range userEvents {
			fmt.Println("Received Event:", msg)
		}
	}()

	bus.Publish("user_created", "User Rohit just signed up!")
	time.Sleep(100 * time.Millisecond) // Allow event processing
}