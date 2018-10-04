package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	"golang.org/x/crypto/ed25519"
)

func create_group(CoorEndpoint string, transaction *model.Transaction, UserPublicKey ed25519.PublicKey, UserPrivateKey ed25519.PrivateKey) []byte {

	client := cc.GetClient(CoorEndpoint)
	return client.GroupCreate(transaction, UserPublicKey, UserPrivateKey)

}
