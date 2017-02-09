package datastore

import (
	"fmt"
	"github.com/Firemango/octopz/state"
	"gopkg.in/redis.v5"
	"testing"
	"time"
)

func GetDatastore() Datastore {

	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
		Password: "",
		DB:       0,
	})

	return Datastore{
		Context:          &state.Context{Redis: r},
		BroadcastChannel: make(chan Message)}
}

func TestBroadcastMessage(t *testing.T) {
	ds := GetDatastore()
	rc := make(chan Message)
	run := false

	message := Message{
		Key:   "test",
		Value: "test"}

	go ds.Start()
	go ds.Broadcast(message)

	go func() {
		select {
		case msg := <-rc:
			run = true
			if msg.Value != "test" {
				t.Error(fmt.Sprintf("Wrong value from listener: %s", msg))
			}
		}
	}()

	go ds.Listen("test", rc)

	time.Sleep(100 * time.Millisecond)
	if run != true {
		t.Error("Test timed out")
	}
}
