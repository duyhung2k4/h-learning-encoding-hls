package service

import (
	"app/config"
	"app/constant"
	"app/dto/queuepayload"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type encodingService struct{}

type EncodingService interface {
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

func NewEncodingService() EncodingService {
	return &encodingService{}
}
