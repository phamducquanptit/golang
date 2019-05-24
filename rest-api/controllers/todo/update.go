package todo

import (
	. "rest-api/controllers"
	"rest-api/models"
	"rest-api/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router /todo/{id}	[put]
// @Tags Todo
// @Summary Update todo
// @Description Update todo
// @Accept json
// @Produce json
// @Param id path string true "id todo update"
// @Param body body todo.UpdateRequest true "data update"
// @Success 200 string string "{"code": 0, "message": "OK", "data": "{"todo": {"title": "build rest api todo list demo", "status": 0} }"}"
// @Failure 400 string string errno.BadRequest
// @Failure 404 string string errno.NotFound
// @Security OAuth2Application
func Update(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("userID").(uint64)

	var r UpdateRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	td := models.TodoModel{
		Title:  r.Title,
		Status: uint64(r.Status),
	}

	if err := td.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	td.ID = uint64(todoID)
	if err := td.Update(userID); err != nil {
		SendResponse(c, errno.ErrUpdateTodo, nil)
		return
	}
	SendResponse(c, errno.OK, UpdateResponse{
		Todo: r,
	})
}
