package main

import (
	"github.com/AlekseySauron/taskproc/pkg/actions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("", actions.GettingWithoutParam)
	router.GET("/:param", actions.GettingWithParam)
	router.POST("/*param", actions.Posting)
	router.PUT("/", actions.Putting)
	router.DELETE("", actions.Deleting)

	router.Run()
}
