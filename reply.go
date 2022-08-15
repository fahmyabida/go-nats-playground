package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func reply(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("replier", id)
	// Subscribe
	if _, err := nc.Subscribe("REPLY.received", func(m *nats.Msg) {
		fmt.Println(string(m.Data))

		response := []byte(fmt.Sprint(id, ". yes mam, now is ", time.Now()))
		fmt.Println(string(response))
		m.Respond(response)
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
