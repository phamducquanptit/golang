package user

import (
	. "rest-api/controllers"
	"rest-api/models"
	"rest-api/pkg/auth"
	"rest-api/pkg/errno"
	"rest-api/pkg/token"

	"github.com/gin-gonic/gin"
)

// @Router /auth/login [post]
// @Tags Auth
// @Summary Login user
// @Description Login user
// @Produce  json
// @Param user body user.SwaggerLoginBody true "body: username, password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"user": {"id": 1,"created_at": "2019-05-15T16:44:54+07:00","username": "ebisu","password": "$2a$10$UEC866tSde0fZTyRQBTOvOHv9T4qUMgUBhREXkeGWG5s2gHXXWCzm"}, "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Failure 400 string string errno.BadRequest
// @Failure 404 string string errno.NotFound
func Login(c *gin.Context) {
	var u models.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, check := models.FindUser(u.Username)
	if !check {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrTokenInvalid, nil)
		return
	}
	SendResponse(c, errno.OK, LoginResponse{
		User:  *d,
		Token: t,
	})
}
