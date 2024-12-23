package encodingservice

import (
	queuepayload "app/internal/dto/queue_payload"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (s *encodingService) DownloadFileMp4(payload queuepayload.QueueMp4QuantityPayload) error {
	url := fmt.Sprintf("%s/%s", payload.IpServer, payload.Path)

	log.Println("URL MP4: ", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	filepathSave := fmt.Sprintf("cmd/encoding-service/data/video/%s", payload.Path)
	out, err := os.Create(filepathSave)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
