package crypt0cloud_sdk_go

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/crypt0cloud/core/model"
	"github.com/onlyangel/apihandlers"
	"golang.org/x/crypto/ed25519"
)

func (c Crypt0Client) PostSingleTransaction(endpoint string, transaction *model.Transaction, publicKey, privateKey []byte) []byte {
	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	transaction.Content = base64.StdEncoding.EncodeToString(jsonstr)

	sha_256 := sha256.New()
	sha_256.Write(jsonstr)
	contentsha := sha_256.Sum(nil)
	transaction.Hash = base64.StdEncoding.EncodeToString(contentsha)

	sign := ed25519.Sign(privateKey, contentsha)
	transaction.Sign = base64.StdEncoding.EncodeToString(sign)

	transaction.Signer = base64.StdEncoding.EncodeToString(publicKey)

	jsonstr, err = json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	return c._post("http://"+endpoint+"/api/v1/post_single_transaction", jsonstr)
}
