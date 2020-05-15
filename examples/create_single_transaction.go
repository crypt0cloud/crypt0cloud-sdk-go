package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
	"golang.org/x/crypto/ed25519"
)

func create_single_transaction(Endpoint string, transaction *model.Transaction, UserPublicKey ed25519.PublicKey, UserPrivateKey ed25519.PrivateKey) []byte {

	client := cc.GetClient(Endpoint)
	return client.PostSingleTransaction(transaction, UserPublicKey, UserPrivateKey)

}
