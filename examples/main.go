package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"os"
)

var endpoint1 string
var endpoint2 string
var coordinator_endpoint string

var MKPublicKey ed25519.PublicKey
var MKPrivateKey ed25519.PrivateKey

func main(){
	flag.StringVar(&endpoint1, "endpoint 1", "localhost:8081", "url of the endpoint 1")
	flag.StringVar(&endpoint2, "endpoint 2", "localhost:8080", "url of the endpoint 2")
	flag.StringVar(&coordinator_endpoint, "coordinator", "localhost:8080", "url of the coordinator endpoint")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Initin coorinator\n")
	reader.ReadString('\n')

	MKPublicKey, _ := get_key_pair()

	coordinator_init(coordinator_endpoint, MKPublicKey)
}



