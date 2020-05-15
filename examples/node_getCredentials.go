package main

import (
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
)

func node_get_credentials(NodeEndpoint string) *model.NodeIdentification {

	client := cc.GetClient(NodeEndpoint)
	return client.Node_GetCredentials()

}
