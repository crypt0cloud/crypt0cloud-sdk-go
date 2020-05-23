package main

import (
	model "github.com/crypt0cloud/model_go"

	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func get_signing_request(Endpoint string, transaction_id int64) *model.Transaction {

	client := cc.GetClient(Endpoint)
	return client.SigningRequestGet(transaction_id)

}
