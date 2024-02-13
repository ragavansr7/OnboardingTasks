package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// func NewMessage() Message {
// 	return Message("Hi there!")
// }

func NewMessage(phrase string) Message {
    return Message(phrase)
}

// func NewGreeter(m Message) Greeter {
// 	return Greeter{Message: m}
// }

func NewGreeter(m Message) Greeter {
    var grumpy bool
    if time.Now().Unix()%2 == 0 {
        grumpy = true
    }
    return Greeter{Message: m, Grumpy: grumpy}
}
// func (g Greeter) Greet() Message {
// 	return g.Message
// }
func (g Greeter) Greet() Message {
    if g.Grumpy {
        return Message("Go away!")
    }
    return g.Message
}

// func NewEvent(g Greeter) Event {
// 	return Event{Greeter: g}
// }

func NewEvent(g Greeter) (Event, error) {
    if g.Grumpy {
        return Event{}, errors.New("could not create event: event greeter is grumpy")
    }
    return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// func main() {
// 	// fmt.Println(NewMessage())
// 	// fmt.Println(NewGreeter("rags"))
// 	message := NewMessage()
// 	fmt.Println("1" ,message)
//     greeter := NewGreeter(message)
// 	fmt.Println("2" ,greeter)
//     event := NewEvent(greeter)
// 	fmt.Println("3" ,event)
// 	fmt.Println("1111")
//     event.Start()
// 	fmt.Println("1111")

// }
// func main() {
//     e := InitializeEvent()

//     e.Start()
// }

func main() {
    e, err := InitializeEvent("hey ")
    if err != nil {
        fmt.Printf("failed to create event: %s\n", err)
        os.Exit(2)
    }
    e.Start()
}
