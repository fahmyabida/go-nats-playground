package main

import (
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var (
	nc  *nats.Conn
	js  nats.JetStreamContext
	err error
)

func init() {
	nc, err = nats.Connect("localhost:4222")
	if err != nil {
		panic(err)
	}
	// js, err = nc.JetStream(nats.PublishAsyncMaxPending(256))
	// if err != nil {
	// 	panic(err)
	// }
}

func main() {
	cobraCmd := &cobra.Command{Use: "go-nats-playground"}

	cobraCmd.AddCommand(&cobra.Command{
		Use: "subscribe",
		Run: subscribe,
	})
	cobraCmd.AddCommand(&cobra.Command{
		Use: "publish",
		Run: publish,
	})
	cobraCmd.AddCommand(&cobra.Command{
		Use: "reply",
		Run: reply,
	})
	cobraCmd.AddCommand(&cobra.Command{
		Use: "reply-queue",
		Run: replyQueue,
	})
	cobraCmd.AddCommand(&cobra.Command{
		Use: "request",
		Run: request,
	})
	cobraCmd.PersistentFlags().String("id", "", "identifier")
	cobraCmd.Flags()
	if err := cobraCmd.Execute(); err != nil {
		panic(err)
	}
}
