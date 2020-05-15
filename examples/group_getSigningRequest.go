package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
)

func get_signing_request(Endpoint string, transaction_id int64) *model.Transaction {

	client := cc.GetClient(Endpoint)
	return client.SigningRequestGet(transaction_id)

}
