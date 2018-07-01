package user

import (
	"strconv"
	
	"github.com/gin-gonic/gin"
	. "github.com/sivdead/apiserver/handler"
	"github.com/sivdead/apiserver/model"
	"github.com/sivdead/apiserver/pkg/errno"
)

func Delete(c *gin.Context)  {
	userId, _ := strconv.Atoi(c.Params.ByName("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	
	SendResponse(c, nil, nil)
}