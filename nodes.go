package crypt0cloud_sdk_go

import (
	"encoding/json"
	"github.com/crypt0cloud/core/model"
	"github.com/onlyangel/apihandlers"
)

func (c Crypt0Client) GetRemoteNodeCredentials(endpoint string) *model.NodeIdentification {
	//TODO: CHANGE URL WHEN BLOCK CHANGES
	response := c._get("http://" + endpoint + "/api/v1/node_id")

	nodeI := new(model.NodeIdentification)
	err := json.Unmarshal(response, nodeI)
	apihandlers.PanicIfNotNil(err)

	return nodeI
}
