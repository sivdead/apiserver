package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/sivdead/apiserver/handler"
	"github.com/sivdead/apiserver/model"
	"github.com/sivdead/apiserver/pkg/errno"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Params.ByName("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	
	SendResponse(c, nil, user)
}