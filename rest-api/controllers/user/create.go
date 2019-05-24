package user

import (
	_log "log"
	. "rest-api/controllers"
	"rest-api/models"
	"rest-api/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Router /user [post]
// @Tags User
// @Summary Create a new user
// @Description Create a new user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"user": { "id": 2,"created_at": "2019-05-15T16:54:05.8135074+07:00","username": "quanpham","password": "$2a$10$Kb4PV4dc.jaUnnjTviWMB.DY0JaHdKBRSEegYrWUZhgXHJQ.tkuu6"}}"
// @Failure 400 string string errno.BadRequest
// @Failure 404 string string errno.NotFound
func Create(c *gin.Context) {

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := models.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		_log.Println("error: ", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, DataResponse{
		User: u,
	})
}
