package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-management/internal/types"
)

func (h *Handler) createTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input types.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Task.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllTasks(c *gin.Context) {

}

func (h *Handler) getTaskById(c *gin.Context) {

}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}
