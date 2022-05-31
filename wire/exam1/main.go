package main

import "fmt"

var msg = "Hello World!"

// Message message type
type Message string

// NewMessage new message
func NewMessage() Message {
	return Message(msg)
}

// People people
type People struct {
	name    string
	message Message
}

// SayHello say hello
func (p People) SayHello() string {
	return fmt.Sprintf("%s 对世界说：%s\n", p.name, p.message)
}

// Event event type
type Event struct {
	people People
}

func (e Event) start() {
	msg := e.people.SayHello()
	fmt.Println(msg)
}

// NewPeople new people
func NewPeople(m Message) People {
	return People{
		name:    "小明",
		message: m,
	}
}

// NewEvent new event
func NewEvent(p People) Event {
	return Event{
		people: p,
	}
}

func main() {
	message := NewMessage()
	people := NewPeople(message)
	event := NewEvent(people)

	event.start()
}
