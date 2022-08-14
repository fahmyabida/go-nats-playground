package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func subscribe(cmd *cobra.Command, args []string) {
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("updates", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
