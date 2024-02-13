// wire.go
package main

import (
	"github.com/google/wire"
)

// func InitializeEvent() Event {
// 	wire.Build(NewEvent, NewGreeter, NewMessage)
// 	return Event{}
// }

// wire.go

// func InitializeEvent() (Event, error) {
//     wire.Build(NewEvent, NewGreeter, NewMessage)
//     return Event{}, nil
// }
func InitializeEvent(phrase string) (Event, error) {
    wire.Build(NewEvent, NewGreeter, NewMessage)
    return Event{}, nil
}
