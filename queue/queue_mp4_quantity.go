package queue

import (
	"app/config"
	"app/dto/queuepayload"
	"app/service"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type queueMp4Quantity struct {
	encodingService service.EncodingService
}
type QueueMp4Quantity interface {
	Worker()
}

func (q *queueMp4Quantity) Worker() {
	queueName := config.GetQueueQuantity()
	conn := config.GetRabbitmq()
	ch, err := conn.Channel()
	if err != nil {
		log.Println("error chanel: ", err)
		return
	}

	qe, err := ch.QueueDeclare(
		string(queueName),
		true,
		false,
		false,
		false,
		amqp091.Table{},
	)
	if err != nil {
		log.Println("error queue declare: ", err)
		return
	}
	log.Printf("start %s", string(queueName))

	msgs, err := ch.Consume(
		qe.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("error consumer: ", err)
		return
	}

	for d := range msgs {
		var payload queuepayload.QueueMp4QuantityPayload

		err := json.Unmarshal(d.Body, &payload)
		if err != nil {
			log.Println("error msg: ", err)
			d.Reject(true)
			continue
		}

		err = q.encodingService.DownloadFileMp4(payload)
		if err != nil {
			log.Println("error download mp4: ", err)
			d.Reject(true)
			continue
		}

		err = q.encodingService.Encoding(payload.Uuid)
		if err != nil {
			log.Println("error encoding hls: ", err)
			d.Reject(true)
			continue
		}

		d.Ack(false)
	}
}

func NewQueueMp4Quantity() QueueMp4Quantity {
	return &queueMp4Quantity{
		encodingService: service.NewEncodingService(),
	}
}
