package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
)

func create_signing_request(Endpoint string, transaction *model.Transaction) []byte {

	client := cc.GetClient(Endpoint)
	return client.SigningRequestCreate(transaction)

}
