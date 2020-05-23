package main

import (
	model "github.com/crypt0cloud/model_go"
	"golang.org/x/crypto/ed25519"

	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func create_app(appname, CoorEndpoint, callback string, MKPublicKey ed25519.PublicKey, MKPrivateKey ed25519.PrivateKey) (*model.Transaction, []byte, []byte) {

	client := cc.GetClient(CoorEndpoint)
	NewAppTransaction, AppPublicKey, AppPrivateKey := client.Coord_CreateAPP(CoorEndpoint, MKPublicKey, MKPrivateKey, callback, appname)

	return NewAppTransaction, AppPublicKey, AppPrivateKey

}
