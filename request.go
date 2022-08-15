package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func request(cmd *cobra.Command, args []string) {
	msg, err := nc.RequestMsg(&nats.Msg{
		Subject: "ORDERS.received",
		Header: map[string][]string{
			"": {},
		},
		Data: []byte("what time is it ?"),
	}, 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Subject)
	// Use the response
	log.Printf("Reply: %s", msg.Data)
}
