package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type mess struct {
	TypeMess string `json:"typeMess"`
	Text     string `json:"text"`
}

// var tasksVol = []*task{
// 	&task{ID: "1", Name: "task1"},
// 	&task{ID: "2", Name: "task2"},
// 	&task{ID: "3", Name: "task3"},
// }

var tasks = []task{
	{ID: "1", Name: "task1"},
	{ID: "2", Name: "task2"},
	{ID: "3", Name: "task3"},
}

func GettingWithoutParam(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func GettingWithParam(c *gin.Context) {
	paramId := c.Param("param")

	for _, curTask := range tasks {

		if curTask.ID == paramId {
			c.JSON(http.StatusOK, curTask)

			return
		}

	}

	mes := mess{"error", "task not found"}
	c.JSON(http.StatusNotFound, mes)
}

func Posting(c *gin.Context) {
	var newTask task

	err := c.BindJSON(&newTask)
	if err != nil {
		return
	}

	for _, curTask := range tasks {
		if curTask.ID == newTask.ID {

			mes := mess{"error", "task id already exists"}
			c.JSON(http.StatusBadRequest, mes)

			return
		}
	}

	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func Putting(c *gin.Context) {
	var inTasks []task

	err := c.BindJSON(&inTasks)

	if err != nil {
		return
	}

	for _, curTask := range inTasks {
		updateName(curTask.ID, curTask.Name)
	}

	c.JSON(http.StatusOK, tasks)

}

func updateName(ID string, Name string) bool {
	for i := 0; i < cap(tasks); i++ {
		curTask := tasks[i]
		if curTask.ID == ID {
			tasks[i].Name = Name
			return true
		}
	}
	return false

}

func Putting2(c *gin.Context) {
	var inTasks []task

	jsonDataBytes, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(jsonDataBytes, &inTasks); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for _, curTask := range inTasks {
		updateName(curTask.ID, curTask.Name)
	}

	c.JSON(http.StatusOK, tasks)

}

func Deleting(c *gin.Context) {
	var inTasks []task

	err := c.BindJSON(&inTasks)

	if err != nil {
		return
	}

	for _, curTask := range inTasks {
		deleteTask(curTask.ID)
	}

	c.JSON(http.StatusOK, tasks)
}

func deleteTask(ID string) {
	//for i := 0; i < cap(tasks); i++ {
	for i := 0; i < len(tasks); i++ {
		curTask := tasks[i]
		if curTask.ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}
