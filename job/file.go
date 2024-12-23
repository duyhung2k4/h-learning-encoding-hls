package jobapp

import (
	"app/internal/connection"
	constant "app/internal/constants"
	"app/internal/entity"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

type fileJob struct {
	psql *gorm.DB
}

type FileJob interface {
	DeleteDirEncoding()
}

func (j *fileJob) DeleteDirEncoding() {
	listDir, err := os.ReadDir("cmd/encoding-service/data/encoding")
	if err != nil {
		log.Println("error get list file: ", err)
		return
	}

	listUuid := []string{}
	for _, f := range listDir {
		if !f.IsDir() {
			continue
		}
		uuid := strings.Split(f.Name(), ".")[0]
		listUuid = append(listUuid, uuid)
	}

	quantity := constant.QUANTITY_MAP[connection.GetConnect().QueueQuantity]
	fieldQuantity := fmt.Sprintf("url%s", quantity.Resolution)

	var listVideoLession []entity.VideoLession
	err = j.psql.
		Model(&entity.VideoLession{}).
		Where(`
			code IN ?
			AND ? IS NOT NULL
		`, listUuid, gorm.Expr(fieldQuantity)).
		Find(&listVideoLession).Error
	if err != nil {
		log.Println("error get listVideoLession: ", err)
		return
	}

	listError := []error{}
	for _, v := range listVideoLession {
		path := fmt.Sprintf("cmd/encoding-service/data/encoding/%s", v.Code)
		err := os.RemoveAll(path)
		if err != nil {
			listError = append(listError, err)
		}
	}

	if len(listError) > 0 {
		for _, e := range listError {
			log.Println("error delete file mp4: ", e)
		}

		return
	}
}

func NewFileJob() FileJob {
	return &fileJob{
		psql: connection.GetPsql(),
	}
}
