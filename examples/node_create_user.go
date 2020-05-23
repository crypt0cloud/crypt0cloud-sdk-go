package main

import (
	model "github.com/crypt0cloud/model_go"

	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func node_create_user(NodeEndpoint string) (*model.Transaction, []byte, []byte) {

	client := cc.GetClient(NodeEndpoint)
	return client.Node_CreateUser()

}
