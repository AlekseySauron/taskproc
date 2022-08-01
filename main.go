package main

import (
	db_mod "github.com/AlekseySauron/taskproc/pkg/db_mod"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db_mod.InitDb()
	defer db.Stop()

	if err != nil {
		return
	}

	err = db_mod.FillDb(db)
	if err != nil {
		return
	}

	// countTasksDB := db.Count()
	// fmt.Println("countTasksDB = ", countTasksDB.Count)

	// records := db.RecordsList().Data
	// fmt.Println("res = ", records)
	// for i := 0; i < len(records); i++ {
	// 	fmt.Println("record = ", string(db.Read(records[i]).Data))
	// }

	// return

	router := gin.Default()

	err = http_mod.Http_actions(router, db)
	if err != nil {
		return
	}

	// router.GET("", actions.GettingWithoutParam)
	// router.GET("/:param", actions.GettingWithParam)
	// router.POST("/*param", actions.Posting)
	// router.PUT("/", actions.Putting)
	// router.DELETE("", actions.Deleting)
	// router.Run()
}
