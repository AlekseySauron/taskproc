package myservices

import (
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

	records := getRecords(h.db) //.db.RecordsList().Data
	c.JSON(http.StatusOK, listData(h.db, records))

}

func (h *Handler) GettingWithParam(c *gin.Context) {
	paramId := c.Param("param")

	taskExists, curTask := getTaskByID(h.db, paramId)
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

		taskExists, _ := getTaskByID(h.db, newTask.ID)
		if !taskExists {
			updateName(h.db, newTask.ID, newTask.Name)
		}
	}

	records := h.db.RecordsList().Data
	c.JSON(http.StatusOK, listData(h.db, records))
}

//только обновляет
func (h *Handler) Putting(c *gin.Context) {
	var inTasks []task

	err := c.BindJSON(&inTasks)

	if err != nil {
		return
	}

	for _, curTask := range inTasks {
		taskExists, _ := getTaskByID(h.db, curTask.ID)
		if taskExists {
			updateName(h.db, curTask.ID, curTask.Name)
		}
	}

	records := h.db.RecordsList().Data
	c.JSON(http.StatusOK, listData(h.db, records))
}

func (h *Handler) Deleting(c *gin.Context) {
	var inTasks []task

	err := c.BindJSON(&inTasks)

	if err != nil {
		return
	}

	for _, curTask := range inTasks {
		deleteTask(h.db, curTask.ID)
	}

	//c.JSON(http.StatusOK, tasks)
	records := getRecords(h.db) //.db.RecordsList().Data
	c.JSON(http.StatusOK, listData(h.db, records))
}
