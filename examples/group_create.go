package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
	"golang.org/x/crypto/ed25519"
)

func create_group(Endpoint string, transaction *model.Transaction, AppPublicKey ed25519.PublicKey, AppPrivateKey ed25519.PrivateKey) []byte {

	client := cc.GetClient(Endpoint)
	return client.GroupCreate(transaction, AppPublicKey, AppPrivateKey)

}
