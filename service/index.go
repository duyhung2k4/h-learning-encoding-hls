package service

import encodingservice "app/cmd/encoding-service/service/encoding"

func Register() Service {
	return Service{
		EncodingService: encodingservice.Register(),
	}
}
