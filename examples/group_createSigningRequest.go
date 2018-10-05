package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func create_signing_request(Endpoint string, transaction *model.Transaction) []byte {

	client := cc.GetClient(Endpoint)
	return client.SigningRequestCreate(transaction)

}
