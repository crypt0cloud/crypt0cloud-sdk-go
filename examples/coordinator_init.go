package main

import (
	"golang.org/x/crypto/ed25519"

	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func coordinator_init(endpoint string, MKPublicKey ed25519.PublicKey) {

	//Create client and register master public key to setup coordinator
	client := cc.GetClient(endpoint)
	client.Coord_RegisterMasterkey(MKPublicKey)

}
