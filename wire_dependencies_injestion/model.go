package main

type Message string

type Greeter struct {
    Message Message // <- adding a Message field
	Grumpy bool
}

type Event struct {
    Greeter Greeter // <- adding a Greeter field
}