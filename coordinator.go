package crypt0cloud_sdk_go

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/onlyangel/apihandlers"
	"golang.org/x/crypto/ed25519"

	crypto "github.com/crypt0cloud/crypto_go"
	model "github.com/crypt0cloud/model_go"
)

func (d Crypt0Client) Coord_RegisterMasterkey(public_key []byte) {
	pkey := crypto.Base64_encode(public_key)

	returned := d._get("http://" + d.Endpoint + "/api/v1/coord/register_masterkey?url=" + d.Endpoint + "&key=" + url.QueryEscape(pkey))

	myerror := new(apihandlers.ErrorType)
	json.Unmarshal(returned, myerror)

	if myerror.Error != "" {
		apihandlers.PanicWithMsg(myerror.Error)
	}
}

func (d Crypt0Client) Coord_AddNode(coordinator_private []byte, coord_endpoint, node_endpoint string) string {
	//ARRAY OF URLS
	data := struct {
		Urls []string
	}{
		[]string{node_endpoint},
	}

	jsonstr, err := json.Marshal(data)
	apihandlers.PanicIfNotNil(err)

	sha_256 := sha256.New()
	sha_256.Write(jsonstr)
	contentsha := sha_256.Sum(nil)
	base64content := base64.StdEncoding.EncodeToString(jsonstr)

	sign := ed25519.Sign(coordinator_private, contentsha)
	base64sign := base64.StdEncoding.EncodeToString(sign)

	tosend := struct {
		Content string
		Sign    string
	}{
		base64content,
		base64sign,
	}

	jsonstr, err = json.Marshal(tosend)
	fmt.Printf("%s\n", string(jsonstr))
	apihandlers.PanicIfNotNil(err)

	returned := d._post("http://"+coord_endpoint+"/api/v1/coord/register_nodes", jsonstr)

	myerror := new(apihandlers.ErrorType)
	json.Unmarshal(returned, myerror)

	if myerror.Error != "" {
		apihandlers.PanicWithMsg(myerror.Error)
	}

	return string(returned)
}

func (c Crypt0Client) Coord_CreateAPP(coord_endpoint string, coord_publ, coord_priv []byte, callback, appname string) (*model.Transaction, []byte, []byte) {
	fmt.Printf("Creating new App\n")

	nodeID := c.Node_GetCredentials()
	//lastNodes := c.Block_getLasts()

	appPublicKey, appPrivateKey, err := ed25519.GenerateKey(rand.New(rand.NewSource(time.Now().UnixNano())))
	apihandlers.PanicIfNotNil(err)

	transaction := new(model.Transaction)
	transaction.SignerKinds = []string{"NewApp"}
	transaction.SignKind = "NewApp"
	transaction.AppID = base64.StdEncoding.EncodeToString(appPublicKey)
	transaction.Parent = ""
	transaction.Callback = callback
	transaction.Payload = appname
	transaction.BlockSign = "trasient"

	transaction.FromNode = *nodeID
	transaction.ToNode = *nodeID

	transaction.Creation = time.Now().UnixNano()

	transaction = signTransaction(transaction, coord_publ, coord_priv)

	jsonstr, err := json.Marshal(transaction)
	apihandlers.PanicIfNotNil(err)

	returned := c._post("http://"+coord_endpoint+"/api/v1/coord/add_app", jsonstr)
	apihandlers.PanicIfNotNil(err)

	myerror := new(apihandlers.ErrorType)
	json.Unmarshal(returned, myerror)

	if myerror.Error != "" {
		apihandlers.PanicWithMsg(myerror.Error)
	}

	return transaction, appPublicKey, appPrivateKey
}
