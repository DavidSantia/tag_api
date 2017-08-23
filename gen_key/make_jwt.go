package main

import (
	"encoding/json"
	"fmt"

	"github.com/DavidSantia/tag_api"
	"github.com/dvsekhvalnov/jose2go"
)

func main() {

	pl := tag_api.JwtPayload{
		UserId: 3,
		Guid: "568e1607-752d-4853-ae21-a1b29d3359f6",
	}
	//pl := tag_api.JwtPayload{
	//	UserId: 10,
	//	Guid: "3f88265e-b2b2-450c-878d-59f14ee76bfd",
	//}

	payload, err := json.Marshal(pl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Payload = %s\n", payload)

	key := tag_api.JwtKey

	token, err := jose.Encrypt(string(payload), jose.A128KW, jose.A128GCM, key)
	if err != nil {
		panic("invalid key format")
	}

	fmt.Printf("\ntoken = %v\n", token)
}
