package crypt0cloud_sdk_go

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/onlyangel/apihandlers"
	"golang.org/x/crypto/ed25519"

	model "github.com/crypt0cloud/model_go"
)

func (c Crypt0Client) GroupCreate(transaction *model.Transaction, publicKey, privateKey []byte) []byte {
	transaction = signTransaction(transaction, publicKey, privateKey)

	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	return c._post("http://"+c.Endpoint+"/api/v1/create_group", jsonstr)
}

func (c Crypt0Client) SigningRequestCreate(transaction *model.Transaction) []byte {
	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	return c._post("http://"+c.Endpoint+"/api/v1/create_signingRequest", jsonstr)
}

func (c Crypt0Client) SigningRequestGet(transactionid int64) *model.Transaction {
	responses := c._get("http://" + c.Endpoint + "/api/v1/get_signingRequest?id=" + string(transactionid))
	transaction := new(model.Transaction)

	err := json.Unmarshal(responses, transaction)
	apihandlers.PanicIfNotNil(err)

	return transaction
}

func (c Crypt0Client) SigningRequestSign(transaction *model.Transaction, publicKey, privateKey []byte) []byte {
	transaction = signTransaction(transaction, publicKey, privateKey)

	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	return c._post("http://"+c.Endpoint+"/api/v1/sign_signingRequest", jsonstr)
}

func signTransaction(transaction *model.Transaction, publicKey, privateKey []byte) *model.Transaction {
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

	return transaction
}
