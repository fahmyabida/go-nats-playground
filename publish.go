package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

func publish(cmd *cobra.Command, args []string) {
	if err := nc.Publish("updates", []byte(fmt.Sprint("hello ", time.Now()))); err != nil {
		log.Fatal(err)
	}
	time.Sleep(100 * time.Millisecond)
}
