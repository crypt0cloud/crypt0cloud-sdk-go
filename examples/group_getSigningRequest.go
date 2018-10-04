package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func get_signing_request(CoorEndpoint string, transaction_id int64) *model.Transaction {

	client := cc.GetClient(CoorEndpoint)
	return client.SigningRequestGet(transaction_id)

}
