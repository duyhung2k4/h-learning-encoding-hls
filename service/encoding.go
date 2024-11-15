package service

import (
	"app/config"
	"app/constant"
	"app/dto/queuepayload"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/rabbitmq/amqp091-go"
)

type encodingService struct {
	connRabbitmq *amqp091.Connection
}

type EncodingService interface {
	SendMessHandleSuccess(payload queuepayload.QueueMp4QuantityPayload) error
	DownloadFileMp4(payload queuepayload.QueueMp4QuantityPayload) error
	Encoding(uuid string) error
}

func (s *encodingService) DownloadFileMp4(payload queuepayload.QueueMp4QuantityPayload) error {
	url := fmt.Sprintf("%s/%s", payload.IpServer, payload.Path)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	filepathSave := fmt.Sprintf("video/%s", payload.Path)
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

func (s *encodingService) Encoding(uuid string) error {
	quantity := constant.QUANTITY_MAP[config.GetQueueQuantity()]
	mp4File := fmt.Sprintf("video/%s.mp4", uuid)

	videoDir := fmt.Sprintf("encoding/%s", uuid)
	err := os.RemoveAll(videoDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(videoDir, os.ModePerm)
	if err != nil {
		return err
	}

	// create dir file encoding
	hlsOutputDir := fmt.Sprintf("encoding/%s", uuid)
	err = os.MkdirAll(hlsOutputDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	outputFile := fmt.Sprintf("encoding/%s/%s_%s.m3u8", uuid, uuid, quantity.Resolution)
	hlsCmd := exec.Command("ffmpeg",
		"-i", mp4File, // Đường dẫn đến file video đã upload uploads/uploaded_video.mp4
		"-vf", fmt.Sprintf("scale=%s", quantity.Scale), // Chỉnh sửa kích thước video
		"-c:a", "aac", // Mã hóa âm thanh
		"-b:a", "128k", // Tốc độ bit âm thanh
		"-c:v", "libx264", // Mã hóa video
		"-preset", "slow", // Cài đặt mã hóa
		"-hls_time", "5",
		"-hls_list_size", "0",
		"-f", "hls", // Định dạng đầu ra
		outputFile, // Đường dẫn đầu ra
	)

	// Chạy lệnh và ghi lại lỗi
	ouput, err := hlsCmd.CombinedOutput()
	if err != nil {
		log.Println("error ffmpeg: ", string(ouput))
		return err
	}

	return nil
}

func (s *encodingService) SendMessHandleSuccess(payload queuepayload.QueueMp4QuantityPayload) error {
	ch, err := s.connRabbitmq.Channel()

	if err != nil {
		return err
	}

	payloadMess := queuepayload.QueueFileM3U8Payload{
		Path:     fmt.Sprintf("encoding/%s", payload.Uuid),
		IpServer: fmt.Sprintf("http://%s:%s/api/v1", config.GetAppHost(), config.GetAppPort()),
		Uuid:     payload.Uuid,
		Quantity: constant.QUANTITY_MAP[config.GetQueueQuantity()].Resolution,
	}

	payloadJsonString, err := json.Marshal(payloadMess)
	if err != nil {
		return err
	}

	ch.PublishWithContext(context.Background(),
		"",
		string(constant.QUEUE_FILE_M3U8),
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payloadJsonString,
		},
	)

	return nil
}

func NewEncodingService() EncodingService {
	return &encodingService{
		connRabbitmq: config.GetRabbitmq(),
	}
}
