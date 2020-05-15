package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
)

func node_create_user(NodeEndpoint string) (*model.Transaction, []byte, []byte) {

	client := cc.GetClient(NodeEndpoint)
	return client.Node_CreateUser()

}
