package controller

import (
	"log"

	"github.com/vitaodemolay/album-system/internal/services"
)

type Controller struct {
	albumService services.IAlbumService
}

func NewController(conString string) *Controller {
	srv, err := services.NewAlbumService(conString)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return &Controller{albumService: srv}
}
