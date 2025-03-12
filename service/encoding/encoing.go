package encodingservice

import (
	"app/internal/connection"
	constant "app/internal/constants"
	"app/internal/entity"
	logapp "app/pkg/log"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (s *encodingService) Encoding(uuid string) error {
	quantity := constant.QUANTITY_MAP[connection.GetConnect().QueueQuantity]
	mp4File := fmt.Sprintf("data/mp4/%s.mp4", uuid)

	videoDir := fmt.Sprintf("data/video/%s/%s", uuid, quantity.Resolution)
	err := os.RemoveAll(videoDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(videoDir, os.ModePerm)
	if err != nil {
		return err
	}

	// create dir file encoding
	hlsOutputDir := fmt.Sprintf("data/video/%s/%s", uuid, quantity.Resolution)
	err = os.MkdirAll(hlsOutputDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	outputFile := fmt.Sprintf("data/video/%s/%s/%s_%s.m3u8", uuid, quantity.Resolution, uuid, quantity.Resolution)
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

	dataUpdate := entity.VideoLession{}

	switch quantity.Resolution {
	case string(constant.QUANTITY_VIDEO_360P):
		dataUpdate.Url360p = &outputFile
	case string(constant.QUANTITY_VIDEO_480P):
		dataUpdate.Url480p = &outputFile
	case string(constant.QUANTITY_VIDEO_720P):
		dataUpdate.Url720p = &outputFile
	case string(constant.QUANTITY_VIDEO_1080P):
		dataUpdate.Url1080p = &outputFile
	}

	err = s.psql.Where("code = ?", uuid).Updates(&dataUpdate).Error
	if err != nil {
		logapp.Logger("encoding-video", err.Error(), constant.ERROR_LOG)
		return err
	}

	return nil
}
