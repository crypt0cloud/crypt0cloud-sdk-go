package crypt0cloud_sdk_go

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/crypt0cloud/core/model"
	"github.com/onlyangel/apihandlers"
	"golang.org/x/crypto/ed25519"
	"math/rand"
	"time"
)

func (c Crypt0Client) Node_GetCredentials() *model.NodeIdentification {
	//TODO: CHANGE URL WHEN BLOCK CHANGES
	response := c._get("http://" + c.Endpoint + "/api/v1/node_id")

	nodeI := new(model.NodeIdentification)
	err := json.Unmarshal(response, nodeI)
	apihandlers.PanicIfNotNil(err)

	return nodeI
}

func (c Crypt0Client) Node_CreateUser(endpoint string) (*model.Transaction, []byte, []byte) {
	fmt.Printf("Creating new User\n")
	nodeID := c.Node_GetCredentials()

	UserPublicKey, UserPrivateKey, err := ed25519.GenerateKey(rand.New(rand.NewSource(time.Now().UnixNano())))
	apihandlers.PanicIfNotNil(err)

	transaction := new(model.Transaction)
	transaction.SignerKinds = []string{"NewUser"}
	transaction.SignKind = "NewUser"
	transaction.AppID = base64.StdEncoding.EncodeToString(UserPublicKey)
	transaction.Parent = ""
	transaction.Callback = "http://localhost:8081"
	transaction.Payload = randomdata.Email()

	transaction.FromNode = *nodeID
	transaction.ToNode = *nodeID
	transaction.Creation = time.Now().UnixNano()

	response := c.PostSingleTransaction(endpoint, transaction, UserPublicKey, UserPrivateKey)

	err = json.Unmarshal(response, transaction)

	return transaction, UserPublicKey, UserPrivateKey
}
