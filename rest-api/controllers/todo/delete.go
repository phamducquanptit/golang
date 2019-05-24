package todo

import (
	. "rest-api/controllers"
	"rest-api/models"
	"rest-api/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router /todo/{id} [delete]
// @Tags Todo
// @Summary Delete a todo
// @Description Delete a todo
// @Accept json
// @Produce json
// @Param id path string true "id todo delete"
// @Success 200 string models.TodoInfo "{"code": 0, "message": "OK", "data": ""}"
// @Failure 400 string string errno.BadRequest
// @Failure 404 string string errno.NotFound
// @Security OAuth2Application
func Delete(c *gin.Context) {

	todoID, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("userID").(uint64)

	if err := models.DeleteTodo(userID, uint64(todoID)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
