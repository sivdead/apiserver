package user

import (
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "github.com/sivdead/apiserver/handler"
	"github.com/sivdead/apiserver/model"
	"github.com/sivdead/apiserver/pkg/errno"
	"github.com/sivdead/apiserver/util"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Params.ByName("id"))
	
	// Binding the user data.
	var u model.UserModel
	if !c.Bind(&u) {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	
	// We update the record based on the user id.
	u.Id = uint64(userId)
	
	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	
	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	
	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	
	SendResponse(c, nil, nil)
}