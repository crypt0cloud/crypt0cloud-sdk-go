package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func create_signing_request(CoorEndpoint string, transaction *model.Transaction) []byte {

	client := cc.GetClient(CoorEndpoint)
	return client.SigningRequestCreate(transaction)

}
