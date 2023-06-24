package main

import (
	"fmt"
	"encoding/hex"
	"os"
	"net"
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

	// "Convert" and parse broadcastip
	broadcastDest := net.ParseIP(broadCastAddr)
	if (broadcastDest == nil) {
		fmt.Println("Error parsing broadcastIp")
		os.Exit(1)
	}


}