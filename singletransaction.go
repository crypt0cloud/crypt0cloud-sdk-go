package crypt0cloud_sdk_go

import (
	"encoding/json"

	"github.com/onlyangel/apihandlers"

	model "github.com/crypt0cloud/model_go"
)

func (c Crypt0Client) PostSingleTransaction(transaction *model.Transaction, publicKey, privateKey []byte) []byte {
	transaction = signTransaction(transaction, publicKey, privateKey)

	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	return c._post("http://"+c.Endpoint+"/api/v1/post_single_transaction", jsonstr)
}
