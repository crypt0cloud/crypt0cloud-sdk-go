package main

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/ed25519"
)

func get_key_pair() (publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) {
	// Generate Key Pair from random data
	publicKey, privateKey, err := ed25519.GenerateKey(rand.New(rand.NewSource(time.Now().UnixNano())))
	if err != nil {
		log.Panic(err)
	}
	return publicKey, privateKey
}
