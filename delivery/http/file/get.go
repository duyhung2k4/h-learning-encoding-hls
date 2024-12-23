package filehandle

import (
	constant "app/internal/constants"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *fileHandle) GetFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	dir := ctx.Param("dir")
	imagePath := fmt.Sprintf("cmd/encoding-service/data/encoding/%s/%s", dir, filename)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		logapp.Logger("read-file", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	ctx.File(imagePath)
}
