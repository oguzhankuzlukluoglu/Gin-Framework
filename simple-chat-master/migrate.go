package main

import (
	"github.com/solnsumei/simple-chat/models"
	"github.com/solnsumei/simple-chat/utils"
)

func runMigrations() {
	if config, err := utils.LoadConfigVars(); err != nil {
		panic("Failed to set config variables")
	} else {
		err = models.RunMigration(config)
		if err != nil {
			panic(err)
		}
	}
}
