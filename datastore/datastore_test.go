package datastore

import (
	"fmt"
	"github.com/Firemango/octopz/config"
	"github.com/kelseyhightower/envconfig"
	"log"
	"testing"
	"time"
)

func TestBroadcastMessage(t *testing.T) {

	c := config.Config{}
	err := envconfig.Process("octopz", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	ds := NewDatastore(c)
	rc := make(chan Message)
	run := false

	message := Message{
		Key:   "test",
		Value: "test",
	}

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
