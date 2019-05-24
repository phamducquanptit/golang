package todo

import (
	. "rest-api/controllers"
	"rest-api/pkg/errno"
	"rest-api/services"

	"github.com/gin-gonic/gin"
)

// @Router /list [post]
// @Tags Todo
// @Summary Get list todo
// @Description Get list todo
// @Accept json
// @Produce json
// @Param body body todo.ListRequest true "List todos"
// @Success 200 {object} todo.SwaggerListResponse "{"code": 0, "message": "OK", "data": { "total": 1,"todos": [{ "id": 1,"title": "learning golang","status": 0,"created_at": "2019-05-15 16:57:29","updated_at": "2019-05-15 16:57:29"}] }}"
// @Failure 400 string string errno.BadRequest
// @Failure 404 string string errno.NotFound
// @Security OAuth2Application
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userID := c.MustGet("userID").(uint64)

	infos, count, err := services.ListTodo(int(userID), r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		Todos: infos,
		Total: count,
	})
}
