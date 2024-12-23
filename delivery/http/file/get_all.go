package filehandle

import (
	constant "app/internal/constants"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *fileHandle) GetAllFile(ctx *gin.Context) {
	dir := ctx.Param("dir")
	videoDir := fmt.Sprintf("cmd/encoding-service/data/encoding/%s", dir)

	files, err := os.ReadDir(videoDir)
	if err != nil {
		logapp.Logger("read-file", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"files": fileNames,
	})
}
