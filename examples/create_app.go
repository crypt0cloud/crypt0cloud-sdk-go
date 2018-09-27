package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	"golang.org/x/crypto/ed25519"
)

func create_app(CoorEndpoint string, MKPublicKey ed25519.PublicKey, MKPrivateKey ed25519.PrivateKey) (*model.Transaction, []byte, []byte) {

	client := cc.GetClient(CoorEndpoint)
	NewAppTransaction, AppPublicKey, AppPrivateKey := client.Coord_CreateAPP(CoorEndpoint, MKPublicKey, MKPrivateKey)

	return NewAppTransaction, AppPublicKey, AppPrivateKey

}
