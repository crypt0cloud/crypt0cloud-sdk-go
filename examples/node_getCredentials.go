package main

import (
	model "github.com/crypt0cloud/model_go"

	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func node_get_credentials(NodeEndpoint string) *model.NodeIdentification {

	client := cc.GetClient(NodeEndpoint)
	return client.Node_GetCredentials()

}
