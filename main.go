package main

import (
	"context"
	"fmt"

	"github.com/nbd-wtf/go-nostr"
)

const relayUrl string = "wss://relay.damus.io"

func main() {

	relay, err := nostr.RelayConnect(context.Background(), relayUrl)
	if err != nil {
		panic(err)
	}

	filters := []nostr.Filter{{
		Kinds: []int{30023},
		Limit: 15,
	}}

	ctx, cancel := context.WithCancel(context.Background())
	sub, err := relay.Subscribe(ctx, filters)

	go func() {
		<-sub.EndOfStoredEvents
		// handle end of stored events (EOSE, see NIP-15)
		cancel()
	}()

	if err != nil {
		fmt.Print(err)
	}

	var events []*nostr.Event

	for ev := range sub.Events {
		// handle returned event.
		// channel will stay open until the ctx is cancelled (in this case, by calling cancel())
		events = append(events, ev)

		fmt.Print(events)
	}

}
