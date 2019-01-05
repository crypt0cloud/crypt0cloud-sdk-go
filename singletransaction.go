package crypt0cloud_sdk_go

import (
	"encoding/json"
	"github.com/crypt0cloud/core/model"
	"github.com/onlyangel/apihandlers"
)

func (c Crypt0Client) PostSingleTransaction(transaction *model.Transaction, publicKey, privateKey []byte) []byte {
	transaction = signTransaction(transaction, publicKey, privateKey)

	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	return c._post("https://"+c.Endpoint+"/api/v1/post_single_transaction", jsonstr)
}
