package queue

import (
	"app/cmd/encoding-service/service"
	"app/internal/connection"
	queuepayload "app/internal/dto/queue_payload"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type queueMp4Quantity struct {
	service service.Service
}
type QueueMp4Quantity interface {
	Worker()
}

func (q *queueMp4Quantity) Worker() {
	queueName := connection.GetConnect().QueueQuantity
	conn := connection.GetRabbitmq()
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
			d.Reject(false)
			continue
		}

		err = q.service.EncodingService.Encoding(payload.Uuid)
		if err != nil {
			log.Println("error encoding hls: ", err)
			d.Reject(false)
			continue
		}

		d.Ack(false)
	}
}

func NewQueueMp4Quantity() QueueMp4Quantity {
	return &queueMp4Quantity{
		service: service.Register(),
	}
}
