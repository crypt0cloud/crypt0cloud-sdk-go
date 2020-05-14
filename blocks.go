package crypt0cloud_sdk_go

import (
	"encoding/json"
	"github.com/onlyangel/apihandlers"
	model "source.cloud.google.com/crypt0cloud-app/crypt0cloud/model_go"
)

func (d Crypt0Client) Block_getLasts() []model.Block {

	returned := d._get("http://" + d.Endpoint + "/api/v1/block/get_lasts")

	var arr []model.Block

	err := json.Unmarshal(returned, &arr)
	apihandlers.PanicIfNotNil(err)

	return arr
}
