package job

import (
	"app/constant"
	"log"

	"github.com/robfig/cron/v3"
)

func InitJob() {
	fileJob := NewFileJob()

	c := cron.New()

	log.Println("start background job")

	c.AddFunc(constant.EVERY_30S, fileJob.DeleteDirEncoding)

	c.Start()

	select {}
}