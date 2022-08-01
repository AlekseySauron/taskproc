package db_mod

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlekseySauron/taskproc/pkg/actions"
	"github.com/claygod/coffer"
)

func InitDb() (*coffer.Coffer, error) {
	curDir, _ := os.Getwd()
	// dbDir := path.Join(curDir, "data")
	//dbDir := curDir + "\\data\\"
	dbDir := filepath.Join(curDir, "data")

	db, err, wrn := coffer.Db(dbDir).Create()

	switch {
	case err != nil:
		fmt.Println("Error:", err)
		return nil, err
	case wrn != nil:
		fmt.Println("Warning:", err)
		return nil, err
	}
	if !db.Start() {
		fmt.Println("Error: not start")
		err = errors.New("error: not start")
		return nil, err
	}
	//defer db.Stop()
	return db, nil
}

func FillDb(db *coffer.Coffer) error {
	tasks := actions.GetTasks()
	for i := 0; i < len(tasks); i++ {
		curTask := tasks[i]

		// if rep := db.Write(curTask.ID, []byte(curTask.Name)); rep.IsCodeError() {
		// 	fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
		// 	return true
		// }

		rep := db.Write(curTask.ID, []byte(curTask.Name))
		if rep.IsCodeError() {
			fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
			return rep.Error
		}

	}
	return nil
}
