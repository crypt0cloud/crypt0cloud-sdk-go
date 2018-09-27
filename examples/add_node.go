package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	"golang.org/x/crypto/ed25519"
)

func add_node(CoorEndpoint, NodeEndpoint string, MKPrivateKey ed25519.PrivateKey) {

	//Create client and register master public key to setup coordinator
	client := cc.GetClient(CoorEndpoint)
	client.Coord_AddNode(MKPrivateKey, NodeEndpoint)

}
