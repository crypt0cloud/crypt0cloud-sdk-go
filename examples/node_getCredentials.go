package main

import (
	"github.com/crypt0cloud/core/model"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
)

func node_get_credentials(NodeEndpoint string) *model.NodeIdentification {

	//Create client and register master public key to setup coordinator
	client := cc.GetClient(NodeEndpoint)
	return client.Node_GetCredentials()

}
