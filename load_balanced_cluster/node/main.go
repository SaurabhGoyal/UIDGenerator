package main

import (
	"log"
)

func main() {
	log.Printf("In main")
	config := loadConfig()
	generator := initGenerator(config.NodeId)
	webapp := initWebApp(generator)
	registerNodeToLB(config)
	runWebApp(webapp)
}
