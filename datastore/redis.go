package datastore

import (
	"github.com/Firemango/octopz/state"
)

type Message struct {
	Key   string
	Value string
}

type Datastore struct {
	Context          *state.Context
	BroadcastChannel chan Message
}

func (ds *Datastore) Broadcast(message Message) {
	ds.BroadcastChannel <- message
}

func (ds *Datastore) Listen(topic string, c chan Message) {
	sub, err := ds.Context.Redis.Subscribe(topic)
	if err != nil {
		panic(err)
	}
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		c <- Message{Key: topic, Value: msg.Payload}
	}
}

func (ds *Datastore) Start() {
	for {
		message := <-ds.BroadcastChannel

		err := ds.Context.Redis.Set(message.Key, message.Value, 0).Err()
		if err != nil {
			panic(err)
		}

		err = ds.Context.Redis.Publish("test", message.Value).Err()
		if err != nil {
			panic(err)
		}

	}
}
