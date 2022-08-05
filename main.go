package main

import (
	"github.com/AlekseySauron/taskproc/pkg/httpmod"
)

func main() {
	// defer db.Db.Stop()
	router := httpmod.New()

	err := router.Start()
	if err != nil {
		return
	}
}
