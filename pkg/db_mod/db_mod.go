package db_mod

import (
	"errors"
	"fmt"
	"os"

	"github.com/claygod/coffer"
)

type DbObject struct {
	Db *coffer.Coffer
}

type task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var tasks = []task{
	{ID: "10", Name: "task1"},
	{ID: "20", Name: "task2"},
	{ID: "30", Name: "task3"},
	{ID: "40", Name: "task4"},
}

func New() *DbObject {
	res := &DbObject{}

	curDir, _ := os.Getwd()
	//dbDir := filepath.Join(curDir, "data")
	dbDir := curDir + "\\data\\"

	var err, wrn error
	res.Db, err, wrn = coffer.Db(dbDir).Create()

	switch {
	case err != nil:
		fmt.Println("Error:", err)
		panic(err)
	case wrn != nil:
		fmt.Println("Warning:", err)
		panic(err)
	}

	if !res.Db.Start() {
		fmt.Println("Error: not start")
		panic("Ошибка старта БД")
	}

	return res
}

func (dbo *DbObject) Count() int {
	return dbo.Db.Count().Count
}

func (dbo *DbObject) Init() error {
	curDir, _ := os.Getwd()
	//dbDir := filepath.Join(curDir, "data")
	dbDir := curDir + "\\data\\"

	var temp *coffer.Coffer
	temp, err, wrn := coffer.Db(dbDir).Create()
	dbo.Db = temp

	switch {
	case err != nil:
		fmt.Println("Error:", err)
		//return nil, err
		return err
	case wrn != nil:
		fmt.Println("Warning:", err)
		// return nil, err
		return err
	}
	//if !db.Start() {
	if !dbo.Db.Start() {
		fmt.Println("Error: not start")
		err = errors.New("error: not start")
		// return nil, err
		return err
	}
	//defer db.Stop()
	// return db, nil
	return nil

}

func (dbo *DbObject) Fill() error {
	//func Fill(dbo *DbObject) error {
	db := dbo.Db

	var keys []string
	countRec := dbo.Db.Count().Count
	for i := 0; i < countRec; i++ {
		keys = append(keys, fmt.Sprint(i))
	}
	dbo.Db.DeleteListOptional(keys)

	tasks := GetTasks()
	for i := 0; i < len(tasks); i++ {
		curTask := tasks[i]

		rep := db.Write(curTask.ID, []byte(curTask.Name))
		if rep.IsCodeError() {
			fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
			return rep.Error
		}

	}
	return nil

}

func GetTasks() []task {
	return tasks
}
