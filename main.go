package main

import (
	"fmt"
	"encoding/hex"
	"os"
)

func main() {

	// Define macAddr, broadcastAddr to send network package to
	const macAddr       = "6C1B9590FE32"
	const broadCastAddr = "255.255.255.255"

	// Convert macAddr in Bytes
	mac, err := hex.DecodeString(macAddr)
	if (err != nil) {
		fmt.Println("Error converting (string) macAddr to byte representation")
		os.Exit(1)
	}

	fmt.Printf("%v", mac) // FIXME

}