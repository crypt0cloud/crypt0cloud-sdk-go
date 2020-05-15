package crypt0cloud

import (
	"encoding/json"
	"math/rand"
	"time"

	"golang.org/x/crypto/ed25519"

	"github.com/crypt0cloud/crypt0cloud-sdk-go"
	crypto "github.com/crypt0cloud/crypto_go"
	model "github.com/crypt0cloud/model_go"
)

type Crypt0Cloud struct {
	Client *crypt0cloud_sdk_go.Crypt0Client
}

func (cc Crypt0Cloud) CreateKey(payload string) (*model.Transaction, ed25519.PublicKey, ed25519.PrivateKey, error) {
	b1 := cc.Client.Block_getLasts()[0]
	nid := cc.Client.Node_GetCredentials()

	publicKey, privateKey, err := ed25519.GenerateKey(rand.New(rand.NewSource(time.Now().UnixNano())))
	if err != nil {
		return nil, nil, nil, err
	}

	t := &model.Transaction{
		AppID:       crypto.Base64_encode(publicKey),
		Payload:     payload,
		SignKind:    "NewUser",
		SignerKinds: []string{"NewUser"},
		FromNode:    *nid,
		ToNode:      *nid,

		Callback:  "demo",
		Creation:  time.Now().UnixNano(),
		BlockSign: b1.Sign,
	}
	t.Signer = t.AppID

	response := cc.Client.PostSingleTransaction(t, publicKey, privateKey)
	err = json.Unmarshal(response, t)
	if err != nil {
		return nil, nil, nil, err
	}

	return t, publicKey, privateKey, nil
}

func (cc Crypt0Cloud) CreateGroup(payload string, kinds []string, appkey_public ed25519.PublicKey, appkey_private ed25519.PrivateKey, callback string) (*model.Transaction, error) {
	nid := cc.Client.Node_GetCredentials()

	t := &model.Transaction{
		AppID:       crypto.Base64_encode(appkey_public),
		Payload:     payload,
		SignKind:    "__NEWCONTRACT",
		SignerKinds: kinds,

		FromNode: *nid,
		ToNode:   *nid,

		Callback: callback,
		Creation: time.Now().UnixNano(),
	}
	t.Signer = t.AppID

	response := cc.Client.GroupCreate(t, appkey_public, appkey_private)

	err := json.Unmarshal(response, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (cc Crypt0Cloud) CreateSigningRequest(payload, signkind string, kinds []string, parent string, appkey_public ed25519.PublicKey, callback string) (*model.Transaction, error) {
	nid := cc.Client.Node_GetCredentials()

	t := &model.Transaction{
		AppID:       crypto.Base64_encode(appkey_public),
		SignKind:    signkind,
		SignerKinds: kinds,
		Parent:      parent,
		FromNode:    *nid,
		Callback:    callback,
		Payload:     payload,
	}

	response := cc.Client.SigningRequestCreate(t)

	err := json.Unmarshal(response, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (cc Crypt0Cloud) GetSigningRequest(transactioid int64) *model.Transaction {
	return cc.Client.SigningRequestGet(transactioid)
}

func (cc Crypt0Cloud) SignSigningRequest(t *model.Transaction, signkey_public ed25519.PublicKey, signkey_private ed25519.PrivateKey) (*model.Transaction, error) {
	b1 := cc.Client.Block_getLasts()[0]

	t.Signer = crypto.Base64_encode(signkey_public)
	t.ToNode = t.FromNode
	t.BlockSign = b1.Sign

	response := cc.Client.SigningRequestSign(t, signkey_public, signkey_private)

	err := json.Unmarshal(response, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
