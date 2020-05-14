package crypt0cloud_sdk_go

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/onlyangel/apihandlers"

	model "github.com/crypt0cloud/model_go"
)

func (c Crypt0Client) Node_SetUp() bool {
	response := c._get("http://" + c.Endpoint + "/api/setup/configure_endpoint?endpoint=" + c.Endpoint)

	return string(response) == "OK"
}

func (c Crypt0Client) Node_GetCredentials() *model.NodeIdentification {
	//TODO: CHANGE URL WHEN BLOCK CHANGES
	response := c._get("http://" + c.Endpoint + "/api/v1/node_id")

	nodeI := new(model.NodeIdentification)
	err := json.Unmarshal(response, nodeI)
	apihandlers.PanicIfNotNil(err)

	return nodeI
}

func (c Crypt0Client) Node_CreateUser(UserPublicKey, UserPrivateKey []byte) (*model.Transaction, []byte, []byte) {
	fmt.Printf("Creating new User\n")
	nodeID := c.Node_GetCredentials()

	transaction := new(model.Transaction)
	transaction.SignerKinds = []string{"NewUser"}
	transaction.SignKind = "NewUser"
	transaction.AppID = base64.StdEncoding.EncodeToString(UserPublicKey)
	transaction.Parent = ""
	transaction.Callback = "http://localhost:8081"
	transaction.Payload = transaction.AppID

	transaction.FromNode = *nodeID
	transaction.ToNode = *nodeID
	transaction.Creation = time.Now().UnixNano()

	response := c.PostSingleTransaction(transaction, UserPublicKey, UserPrivateKey)

	json.Unmarshal(response, transaction)

	return transaction, UserPublicKey, UserPrivateKey
}
