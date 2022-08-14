package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func replyQueue(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("replier-queue ", id)
	// Subscribe
	if _, err := nc.QueueSubscribe("times", "queue-sub", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
		fmt.Println(string(m.Reply))

		response := []byte(fmt.Sprint(id, ". now is ", time.Now()))
		fmt.Println(string(response))
		m.Respond(response)
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
