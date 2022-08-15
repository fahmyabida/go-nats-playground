package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

func publishJs(cmd *cobra.Command, args []string) {
	pub, err := js.Publish("ORDERS.received", []byte(fmt.Sprint("hello ", time.Now())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pub)
	time.Sleep(100 * time.Millisecond)
}
