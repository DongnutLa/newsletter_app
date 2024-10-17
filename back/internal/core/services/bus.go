package services

import (
	"fmt"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
)

// Bus is a ref to bus.Bus
var Bus *bus.Bus

// Monoton is an instance of monoton.Monoton
var Monoton monoton.Monoton

// Init inits the app config
func MessagingInit() {
	// Configure id generator
	node := uint64(1)
	initialTime := uint64(1577865600000)
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}

	// Init an id generator
	var idGenerator bus.Next = m.Next

	// Create a new bus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	// Register topics in here
	b.RegisterTopics(domain.TopicList...)

	Bus = b
	Monoton = m

	fmt.Println("Initialized messaging & topics registered")
}
