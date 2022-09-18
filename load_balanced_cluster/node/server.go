package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	uid "github.com/SaurabhGoyal/Snowflake/uid"
)

type NodeRegistrationRequest struct {
	NodeID uint64
	NodeIP string
}

type WebApp struct {
	Generator uid.Generator
}

func (a *WebApp) requestHandler(w http.ResponseWriter, r *http.Request) {
	uid, _ := a.Generator.Get()
	log.Printf("Got req, generated id - %d", uid)
	io.WriteString(w, fmt.Sprint(uid))
}

func registerNodeToLB(config Config) {
	req := NodeRegistrationRequest{
		NodeID: config.NodeId,
		NodeIP: config.Address,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(config.LBIP+"/internal/node", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("Node registered to LB - %v", respBody)
}

func initWebApp(generator uid.Generator) *WebApp {
	webapp := WebApp{Generator: generator}
	log.Printf("Webapp initialised")
	http.HandleFunc("/id", webapp.requestHandler)
	log.Printf("Routing setup")
	return &webapp
}

func runWebApp(webapp *WebApp) {
	err := http.ListenAndServe(":8765", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed.")
	} else if err != nil {
		log.Printf("Error starting server.")
	}
}
