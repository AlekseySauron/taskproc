package routes

import (
	"fmt"
	"net/http"

	"github.com/claygod/coffer"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *coffer.Coffer
}

type mess struct {
	TypeMess string `json:"typeMess"`
	Text     string `json:"text"`
}

type task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewHandler(db_param *coffer.Coffer) *Handler {
	return &Handler{
		db: db_param,
	}
}

func Register(router *gin.Engine, db *coffer.Coffer) {
	h := NewHandler(db)
	router.GET("", h.GettingWithoutParam)
	router.GET("/:param", h.GettingWithParam)
	router.POST("/*param", h.Posting)
	router.PUT("", h.Putting)
	router.DELETE("", h.Deleting)

}

func (h *Handler) GettingWithoutParam(c *gin.Context) {

	fmt.Println("countTasksDB = ", h.db.Count().Count)

	records := h.db.RecordsList().Data
	c.JSON(http.StatusOK, h.listData(records))

}

func (h *Handler) listData(m []string) []task {
	var curTasks []task

	for i := 0; i < len(m); i++ {
		curId := m[i]
		_, curTask := h.getTaskByID(curId)
		curTasks = append(curTasks, curTask)
	}

	return curTasks
}

func (h *Handler) GettingWithParam(c *gin.Context) {
	paramId := c.Param("param")

	taskExists, curTask := h.getTaskByID(paramId)
	if taskExists {
		c.JSON(http.StatusOK, curTask)
	} else {
		mes := mess{"error", "task not found"}
		c.JSON(http.StatusNotFound, mes)
	}

}

//только создает задачу
func (h *Handler) Posting(c *gin.Context) {
	var newTasks []task

	err := c.BindJSON(&newTasks)
	if err != nil {
		return
	}

	for i := 0; i < len(newTasks); i++ {
		newTask := newTasks[i]

		taskExists, _ := h.getTaskByID(newTask.ID)
		if !taskExists {
			h.updateName(newTask.ID, newTask.Name)
		}
	}

	records := h.db.RecordsList().Data
	c.JSON(http.StatusOK, h.listData(records))
}

//только обновляет
func (h *Handler) Putting(c *gin.Context) {
	var inTasks []task

	err := c.BindJSON(&inTasks)

	if err != nil {
		return
	}

	for _, curTask := range inTasks {
		taskExists, _ := h.getTaskByID(curTask.ID)
		if taskExists {
			h.updateName(curTask.ID, curTask.Name)
		}
	}

	records := h.db.RecordsList().Data
	c.JSON(http.StatusOK, h.listData(records))
}

func (h *Handler) updateName(ID string, Name string) bool {

	if rep := h.db.Write(ID, []byte(Name)); rep.IsCodeError() {
		fmt.Printf("Write error: code `%d` msg `%s`", rep.Code, rep.Error)
		return false
	}

	return true
}

func (h *Handler) Deleting(c *gin.Context) {
	var inTasks []task

	err := c.BindJSON(&inTasks)

	if err != nil {
		return
	}

	for _, curTask := range inTasks {
		h.deleteTask(curTask.ID)
	}

	//c.JSON(http.StatusOK, tasks)
	records := h.db.RecordsList().Data
	c.JSON(http.StatusOK, h.listData(records))
}

func (h *Handler) deleteTask(ID string) {
	h.db.Delete(ID)
}

func (h *Handler) getTaskByID(ID string) (bool, task) {
	var emptyTask task

	rep := h.db.Read(ID)
	if rep.IsCodeError() {
		fmt.Printf("Read error: code `%v` msg `%v`", rep.Code, rep.Error)
		return false, emptyTask
	}
	return true, task{ID, string(rep.Data)}
}
