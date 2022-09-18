package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type WebApp struct {
	Nodes []string
}

type NodeRegistrationRequest struct {
	NodeID uint64
	NodeIP string
}

func (a *WebApp) nodeRegistrationRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req NodeRegistrationRequest
	d := json.NewDecoder(r.Body).Decode(&req)
	log.Printf("Node registration request - %v", d)
	a.Nodes = append(a.Nodes, req.NodeIP)
	for i, node := range a.Nodes {
		log.Printf("Nodes - %d - %v", i, node)
	}
	io.WriteString(w, fmt.Sprintf("Registered successfully, Current Nodes - %v", a.Nodes))
}

func (a *WebApp) idRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got id req at LB.")
	if len(a.Nodes) < 1 {
		io.WriteString(w, "No node available")
		return
	} else {
		rand.Seed(time.Now().Unix())
		nodeIndex := rand.Intn(len(a.Nodes))
		node := a.Nodes[nodeIndex]
		resp, err := http.Get(node + "/id")
		if err != nil {
			panic(err)
		} else {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			io.WriteString(w, string(respBody))
			return
		}
	}
}

func initWebapp() *WebApp {
	webapp := WebApp{Nodes: []string{}}
	log.Printf("LB initialised")
	http.HandleFunc("/id", webapp.idRequestHandler)
	http.HandleFunc("/internal/node", webapp.nodeRegistrationRequestHandler)
	log.Printf("Routing setup")
	return &webapp
}

func runWebapp() {
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed.")
	} else if err != nil {
		log.Printf("Error starting server.")
	}
}
