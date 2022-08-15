package main

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func consumer(cmd *cobra.Command, args []string) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// sub, err := js.Subscribe("ORDERS.received", func(msg *nats.Msg) {
	// 	fmt.Println(msg)
	// 	fmt.Println(string(msg.Data))
	// }, nats.Durable("NEW2"))
	// fmt.Println(sub)
	// fmt.Println(err)

	// pull based
	// js.Subscribe("ORDERS.received", func(msg *nats.Msg) {
	// 	fmt.Println(msg)
	// 	fmt.Println(string(msg.Data))

	// 	msg.Ack()
	// }, nats.DeliverSubject("aaa.ORDERS"), nats.Durable("AAA"), nats.AckNone())

	js.QueueSubscribe("ORDERS.received", "BBB", func(msg *nats.Msg) {
		fmt.Println(msg)
		fmt.Println(string(msg.Data))

		msg.Ack()
	}, nats.DeliverSubject("bbb.ORDERS"), nats.AckNone())

	// fmt.Println(sub)
	// fmt.Println(err)

	// sub, err := js.QueueSubscribeSync("ORDERS.received", "hallo")
	// if err != nil {
	// 	panic(err)
	// }
	// m, err := sub.NextMsg(10 * time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(m)
	// fmt.Println(string(m.Data))

	// sub, err := js.PullSubscribe("ORDERS.received", "NEW")
	// if err != nil {
	// 	panic(err)
	// }
	// msgs, err := sub.Fetch(10)
	// if err != nil {
	// 	if err == nats.ErrTimeout {
	// 		fmt.Println("timeout")
	// 		return
	// 	}
	// 	panic(err)
	// }
	// for i := range msgs {
	// 	fmt.Println(string(msgs[i].Data))
	// }
	wg.Wait()
}
