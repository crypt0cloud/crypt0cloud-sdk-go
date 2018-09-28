package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func node_create_user(NodeEndpoint string) (*model.Transaction, []byte, []byte) {

	client := cc.GetClient(NodeEndpoint)
	return client.Node_CreateUser()

}
