package main

import (
	"log"

	snowflake "github.com/SaurabhGoyal/Snowflake/snowflake"
	uid "github.com/SaurabhGoyal/Snowflake/uid"
)

func initGenerator(nodeID uint64) uid.Generator {
	defaultConfig, _ := snowflake.InitDefaultGeneratorConfig()
	generator, _ := snowflake.InitGenerator(defaultConfig, nodeID)
	uid, _ := generator.Get()
	log.Printf("Snowflake loaded - %v for Node-Id - %d", uid, nodeID)
	return generator
}
