package controller

import (
	"log"

	"github.com/vitaodemolay/album-system/internal/infrastructure"
	"github.com/vitaodemolay/album-system/internal/services"
)

type Controller struct {
	albumService services.IAlbumService
	publisher    infrastructure.IPublisherKafka
}

func NewController(conString string, kafkaConnection string) *Controller {
	srv, err := services.NewAlbumService(conString)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	pubSub, err := infrastructure.NewPublisherKafka(kafkaConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Controller{albumService: srv, publisher: pubSub}
}
