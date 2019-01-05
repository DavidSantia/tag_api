package tag_api

import (
	"github.com/boltdb/bolt"
	"github.com/nats-io/go-nats"
)

func (bs *BoltService) connectBolt() (err error) {

	Log.Info.Printf("Connecting to %s", bs.settings.boltFile)
	bs.boltDb, err = bolt.Open(bs.settings.boltFile, 0644, nil)
	return
}

func (bs *BoltService) ConnectNATS() (err error) {

	natsUrl := "nats://" + bs.settings.hostNATS + ":" + bs.settings.portNATS

	Log.Info.Printf("Connecting to %s", natsUrl)
	bs.nconn, err = nats.Connect(natsUrl)
	return
}
