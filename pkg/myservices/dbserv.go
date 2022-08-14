package myservices

import (
	"fmt"

	"github.com/claygod/coffer"
)

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

func FillDB(db *coffer.Coffer) error {

	var keys []string
	records := db.RecordsList().Data
	for i := 0; i < len(records); i++ {
		keys = append(keys, records[i])
	}

	db.DeleteListOptional(keys)

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

func getRecords(db *coffer.Coffer) []string {
	return db.RecordsList().Data
}

func listData(db *coffer.Coffer, m []string) []task {
	var curTasks []task

	for i := 0; i < len(m); i++ {
		curId := m[i]
		_, curTask := getTaskByID(db, curId)
		curTasks = append(curTasks, curTask)
	}

	return curTasks
}

func getTaskByID(db *coffer.Coffer, ID string) (bool, task) {
	var emptyTask task

	rep := db.Read(ID)
	if rep.IsCodeError() {
		fmt.Printf("Read error: code `%v` msg `%v`", rep.Code, rep.Error)
		return false, emptyTask
	}
	return true, task{ID, string(rep.Data)}
}

func updateName(db *coffer.Coffer, ID string, Name string) bool {

	if rep := db.Write(ID, []byte(Name)); rep.IsCodeError() {
		fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
		return false
	}

	return true
}

func deleteTask(db *coffer.Coffer, ID string) {
	db.Delete(ID)
}
