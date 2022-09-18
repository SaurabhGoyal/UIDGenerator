package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const nodeIDKey = "NODE_ID"
const portKey = "PORT"
const lbIPKey = "LB_IP"
const addressKey = "ADDR"

type Config struct {
	NodeId  uint64
	Port    string
	LBIP    string
	Address string
}

func loadConfig() Config {
	nodeID := os.Getenv(nodeIDKey)
	nodeIDNum, err := strconv.ParseInt(nodeID, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Node-ID must be an integer - %s", nodeID))
	}
	c := Config{
		NodeId:  uint64(nodeIDNum),
		Port:    os.Getenv(portKey),
		LBIP:    os.Getenv(lbIPKey),
		Address: os.Getenv(addressKey),
	}
	log.Printf("Loaded config - %v", c)
	return c
}
