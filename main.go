package main

import (
	"encoding/hex"
	"os"
	"net"
)

func main() {

	// Define macAddr, broadcastAddr to send network package to
	const macAddr       = "6C1B9590FE32"
	const broadCastAddr = "255.255.255.255"

	// Convert macAddr in Bytes
	macAddrBin, err := hex.DecodeString(macAddr)
	if (err != nil) {
		println("Error converting (string) macAddr to byte representation")
		os.Exit(1)
	}

	// "Convert" and parse broadcastip
	broadcastDest := net.ParseIP(broadCastAddr)
	if (broadcastDest == nil) {
		println("Error parsing broadcastIp")
		os.Exit(1)
	}

	/*
	Create the network packet
	https://datatracker.ietf.org/doc/html/draft-cheshire-edns0-owner-option-01
	o  Sync sequence: 48 binary 1s (i.e. 6 bytes of 0xFF)
	o  Sixteen repetitions of the 48-bit MAC address of the sleeping server's network interface
	o  Optional 32-bit or 48-bit 'password'
	*/

	// Write 0xFF six times as first step (we only write the payload, the rest of the network packet will be created by go net)
	magicNetPacket := make([]byte, 102)
	for i := 0; i < 6; i++ {
		magicNetPacket[i] = 0xFF
	}

	// Add Mac-Address 16 times
	for i := 6; i < 102; i += 6 {
		copy(magicNetPacket[i:i+6], macAddrBin)
	}

	// Open UDP Connection
	connection, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: broadcastDest, Port: 9})
	if (err != nil) {
		println("Error opening UDP connection")
		os.Exit(1)
	}

	connection.Write(magicNetPacket)
	connection.Close()

	println("Sent wol packet to: ", macAddr)
	os.Exit(0)

}
