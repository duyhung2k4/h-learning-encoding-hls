package filehandle

import (
	constant "app/internal/constants"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type fileHandle struct{}

type FileHandle interface {
	GetAllFile(ctx *gin.Context)
	GetFile(ctx *gin.Context)
}

func NewHandle() FileHandle {
	return &fileHandle{}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "encoding/:dir",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.GetAllFile,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "encoding/:dir/:filename",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.GetFile,
	})
}
