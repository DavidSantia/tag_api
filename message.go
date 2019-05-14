package tag_api

import (
	"encoding/json"

	"github.com/nats-io/go-nats"
)

func (bs *BoltService) listenNATSSub() {
	var qMsg QueueMessage
	ch := make(chan *nats.Msg, 64)

	Log.Info.Printf("Subscribing to nats channel %q\n", NSub)
	sub, err := bs.nconn.ChanSubscribe(NSub, ch)
	if err != nil {
		Log.Error.Println(err)
		return
	}
	defer sub.Unsubscribe()

	for {
		msg := <-ch
		err = json.Unmarshal(msg.Data, &qMsg)
		if err != nil {
			Log.Error.Println(err)
		}

		switch qMsg.Command {
		case "update":
			Log.Info.Printf("Content Update: %v\n", msg)
		default:
			Log.Info.Printf("Unrecognized command: %s\n", qMsg.Command)
		}
	}
}
