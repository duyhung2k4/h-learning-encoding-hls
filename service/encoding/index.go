package encodingservice

import (
	"app/internal/connection"
	queuepayload "app/internal/dto/queue_payload"

	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type encodingService struct {
	connRabbitmq *amqp091.Connection
	psql         *gorm.DB
}

type EncodingService interface {
	SendMessHandleSuccess(payload queuepayload.QueueMp4QuantityPayload) error
	DownloadFileMp4(payload queuepayload.QueueMp4QuantityPayload) error
	Encoding(uuid string) error
}

func Register() EncodingService {
	return &encodingService{
		connRabbitmq: connection.GetRabbitmq(),
		psql:         connection.GetPsql(),
	}
}
