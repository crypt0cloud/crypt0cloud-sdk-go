package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	"golang.org/x/crypto/ed25519"
)

func add_node(endpoint string, APP_PublicKey ed25519.PublicKey) {

	//Create client and register master public key to setup coordinator
	client := cc.GetClient(endpoint)
	client.Coord_AddNode(APP_PublicKey, endpoint)

}
