package encodingservice

import (
	"app/internal/connection"
	constant "app/internal/constants"
	queuepayload "app/internal/dto/queue_payload"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func (s *encodingService) SendMessHandleSuccess(payload queuepayload.QueueMp4QuantityPayload) error {
	ch, err := s.connRabbitmq.Channel()

	if err != nil {
		return err
	}

	payloadMess := queuepayload.QueueFileM3U8Payload{
		Path: fmt.Sprintf("encoding/%s", payload.Uuid),
		IpServer: fmt.Sprintf(
			"http://%s:%s/api/v1",
			connection.GetConnect().EncodingService.Host,
			connection.GetConnect().EncodingService.Port,
		),
		Uuid:     payload.Uuid,
		Quantity: constant.QUANTITY_MAP[connection.GetConnect().QueueQuantity].Resolution,
	}

	payloadJsonString, err := json.Marshal(payloadMess)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(),
		"",
		string(constant.QUEUE_FILE_M3U8),
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payloadJsonString,
		},
	)

	if err != nil {
		return err
	}

	fileDelete := fmt.Sprintf("cmd/encoding-service/data/video/%s.mp4", payload.Uuid)
	err = os.RemoveAll(fileDelete)
	if err != nil {
		return err
	}

	return nil
}
