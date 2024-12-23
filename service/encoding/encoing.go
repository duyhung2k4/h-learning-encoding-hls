package encodingservice

import (
	"app/internal/connection"
	constant "app/internal/constants"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (s *encodingService) Encoding(uuid string) error {
	quantity := constant.QUANTITY_MAP[connection.GetConnect().QueueQuantity]
	mp4File := fmt.Sprintf("cmd/encoding-service/data/video/%s.mp4", uuid)

	videoDir := fmt.Sprintf("cmd/encoding-service/data/encoding/%s", uuid)
	err := os.RemoveAll(videoDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(videoDir, os.ModePerm)
	if err != nil {
		return err
	}

	// create dir file encoding
	hlsOutputDir := fmt.Sprintf("cmd/encoding-service/data/encoding/%s", uuid)
	err = os.MkdirAll(hlsOutputDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	outputFile := fmt.Sprintf("cmd/encoding-service/data/encoding/%s/%s_%s.m3u8", uuid, uuid, quantity.Resolution)
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

	os.RemoveAll(mp4File)

	return nil
}
