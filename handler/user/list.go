package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/sivdead/apiserver/handler"
	"github.com/sivdead/apiserver/pkg/errno"
	"github.com/sivdead/apiserver/service"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if !c.Bind(&r) {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	
	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	
	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
