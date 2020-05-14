package main

import (
	"fmt"
	"github.com/crypt0cloud/core/crypto"
	cc "github.com/crypt0cloud/crypt0cloud-sdk-go"
	model "github.com/crypt0cloud/model_go"
	"github.com/go-errors/errors"
	"golang.org/x/crypto/ed25519"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main(){
	for n:= 0; n<10;n++{
		main2()
	}
}

func main2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	appkey := "SqINCfWmw8C2gvzEcZG8R/zHdAcxVRefYuPhOvLfJBs="
	defer func() {
		if err := recover(); err != nil {
			err := errors.Wrap(err, 1)
			fmt.Printf("\nERROR: %v", err)
		}
	}()

	client := cc.GetClient("localhost:8080",nil)
	blocks := client.Block_getLasts()
	node := client.Node_GetCredentials()

	UserPublicKey, UserPrivateKey, err := ed25519.GenerateKey(rand.New(rand.NewSource(time.Now().UnixNano())))
	if err != nil {
		log.Panic(err)
	}

	t := new(model.Transaction)
	t.AppID = crypto.Base64_encode(UserPublicKey)
	t.Signer = t.AppID
	t.Payload = t.AppID
	t.SignKind = "NewUser"
	t.SignerKinds = [] string { t.SignKind }
	t.FromNode = *node
	t.ToNode = *node;

	t.Callback = "demo";
	t.Creation = time.Now().UnixNano()
	t.BlockSign = blocks[0].Sign

	response := client.PostSingleTransaction(t, UserPublicKey, UserPrivateKey)
	fmt.Println(string(response))

	times := 6
	c := make(chan string, times)

	rand.Seed(time.Now().UnixNano())

	for n := 0 ; n < times; n++ {
		go func(c chan string) {
			tr := new(model.Transaction)
			tr.AppID = appkey
			tr.Signer = crypto.Base64_encode(UserPublicKey)
			tr.Payload = fmt.Sprintf("payload %d",rand.Uint64())
			tr.SignKind = ""
			tr.SignerKinds = []string{}

			tr.FromNode = *node
			tr.ToNode = *node

			tr.Callback = "demo"
			tr.Creation = time.Now().UnixNano()
			tr.BlockSign = blocks[0].Sign

			response := client.PostSingleTransaction(tr, UserPublicKey, UserPrivateKey)
			fmt.Println(string(response))
			c <- string(response)
		}(c)
	}

	n := 0
	for {
		select {
		case <-c:
			n++
			fmt.Printf("tick: %d\n",n)
			if n == times{
				return
			}

		}
	}



}
