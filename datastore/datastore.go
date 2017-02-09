package datastore

import (
	"fmt"
	"github.com/Firemango/octopz/config"
	"gopkg.in/redis.v5"
)

type Message struct {
	Key   string
	Value string
}

type Datastore struct {
	Redis            *redis.Client
	BroadcastChannel chan Message
}

func (ds *Datastore) Broadcast(message Message) {
	ds.BroadcastChannel <- message
}

func (ds *Datastore) Listen(topic string, c chan Message) {
	sub, err := ds.Redis.Subscribe(topic)
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

		err := ds.Redis.Set(message.Key, message.Value, 0).Err()
		if err != nil {
			panic(err)
		}

		err = ds.Redis.Publish("test", message.Value).Err()
		if err != nil {
			panic(err)
		}

	}
}

func NewDatastore(c config.Config) Datastore {

	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort),
		Password: c.RedisPassword,
		DB:       0,
	})

	ch := make(chan Message)

	ds := Datastore{
		Redis:            r,
		BroadcastChannel: ch,
	}

	go ds.Start()

	return ds
}
