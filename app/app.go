package app

import (
	"context"
	"fmt"

	"github.com/AlekseySauron/taskproc/pkg/myservices"
	"github.com/claygod/coffer"
	"github.com/gin-gonic/gin"
)

type Application struct {
	gin  *gin.Engine
	ctx  context.Context
	repo *coffer.Coffer
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		gin: gin.Default(),
		ctx: ctx,
	}
}

func (a *Application) Run() {
	err := a.getDB()
	if err != nil {
		panic(err)
	}

	myservices.FillDB(a.repo)

	// routes.Register(a.gin, a.repo)
	myservices.Register(a.gin, a.repo)

	a.gin.Run()
}

func (a *Application) Stop() {
	a.repo.Stop()
}

func (a *Application) getDB() error {
	var err, wrn error
	// res.Db, err, wrn = coffer.Db("../data/").Create()
	a.repo, err, wrn = coffer.Db("../data/").Create()

	switch {
	case err != nil:
		fmt.Println("Error:", err)
		panic(err)
	case wrn != nil:
		fmt.Println("Warning:", err)
		panic(err)
	}

	if !a.repo.Start() {
		fmt.Println("Error: not start")
		panic("Ошибка старта БД")
	}

	return err
}
