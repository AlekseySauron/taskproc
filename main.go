package main

import (
	"github.com/AlekseySauron/taskproc/pkg/actions"
	db_mod "github.com/AlekseySauron/taskproc/pkg/db_mod"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db_mod.InitDb()
	defer db.Stop()

	if err != nil {
		return
	}

	//err = db_mod.
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

	router.GET("", actions.GettingWithoutParam)
	router.GET("/:param", actions.GettingWithParam)
	router.POST("/*param", actions.Posting)
	router.PUT("/", actions.Putting)
	router.DELETE("", actions.Deleting)

	router.Run()
}

// func initDb() (*coffer.Coffer, error) {
// 	curDir, _ := os.Getwd()
// 	// dbDir := path.Join(curDir, "data")
// 	dbDir := curDir + "\\data\\"

// 	db, err, wrn := coffer.Db(dbDir).Create()

// 	switch {
// 	case err != nil:
// 		fmt.Println("Error:", err)
// 		return nil, err
// 	case wrn != nil:
// 		fmt.Println("Warning:", err)
// 		return nil, err
// 	}
// 	if !db.Start() {
// 		fmt.Println("Error: not start")
// 		err = errors.New("error: not start")
// 		return nil, err
// 	}
// 	//defer db.Stop()
// 	return db, nil
// }

// func fillDb(db *coffer.Coffer) error {
// 	tasks := actions.GetTasks()
// 	for i := 0; i < len(tasks); i++ {
// 		curTask := tasks[i]

// 		// if rep := db.Write(curTask.ID, []byte(curTask.Name)); rep.IsCodeError() {
// 		// 	fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
// 		// 	return true
// 		// }

// 		rep := db.Write(curTask.ID, []byte(curTask.Name))
// 		if rep.IsCodeError() {
// 			fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
// 			return rep.Error
// 		}

// 	}
// 	return nil
// }
