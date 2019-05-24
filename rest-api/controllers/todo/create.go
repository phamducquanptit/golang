package todo

import (
	_log "log"
	. "rest-api/controllers"
	"rest-api/models"
	"rest-api/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Router /todo [post]
// @Tags Todo
// @Summary Create a new todo
// @Description Create a new todo
// @Accept  json
// @Produce  json
// @Param todo body todo.CreateRequest true "Create a new todo"
// @Success 200 {object} todo.DataResponse "{"code":0,"message":"OK","data": {"todo": {"id": 1,"created_at": "2019-05-15T16:57:29.4549295+07:00","status": 0,"user_id": 1,"title": "learning golang"} } }"
// @Failure 400 string string errno.BadRequest
// @Failure 404 string string errno.NotFound
// @Security OAuth2Application
func Create(c *gin.Context) {

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	userID := c.MustGet("userID").(uint64)

	td := models.TodoModel{
		Title:  r.Title,
		UserID: userID,
	}

	if err := td.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := td.Create(); err != nil {
		_log.Println("error: ", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, DataResponse{
		Todos: td,
	})
}
