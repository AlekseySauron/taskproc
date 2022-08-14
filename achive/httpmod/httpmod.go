package httpmod

import (
	"github.com/AlekseySauron/taskproc/achive/db_mod"
	"github.com/AlekseySauron/taskproc/achive/routes"
	"github.com/gin-gonic/gin"
)

type HttpObject struct {
	gin *gin.Engine
	dbo db_mod.DbObject
}

func New() *HttpObject {
	res := &HttpObject{}
	res.gin = gin.Default()

	dbo := db_mod.New()
	//defer db.Db.Stop()
	err := dbo.Fill()
	if err != nil {
		panic("Ошибка заполнения БД")
	}

	routes.Register(res.gin, dbo.Db)

	return res
}

func (hto *HttpObject) Start() error {

	hto.gin.Run()

	return nil
}
