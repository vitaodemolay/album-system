package services

import (
	"github.com/vitaodemolay/album-system/internal/infrastructure"
	"github.com/vitaodemolay/album-system/internal/model"
)

type IAlbumService interface {
	AddAlbum(album *model.Album) (int, error)
	GetAllAlbums() ([]*model.Album, error)
	GetAlbumById(id string) (*model.Album, error)
	SearchAlbumsByTitle(title string) ([]*model.Album, error)
	DeleteAlbum(id string) error
}

type albumService struct {
	ctx infrastructure.ISqlDbContext
}

func NewAlbumService(conString string) (*albumService, error) {
	ctx, err := infrastructure.NewSqlDbContext(conString)

	if err != nil {
		return nil, err
	}

	return &albumService{ctx: ctx}, nil
}

func (s *albumService) AddAlbum(album *model.Album) (int, error) {
	return s.ctx.AddAlbum(album)
}

func (s *albumService) GetAllAlbums() ([]*model.Album, error) {
	return s.ctx.GetAlbumListByTitle("")
}

func (s *albumService) GetAlbumById(id string) (*model.Album, error) {
	return s.ctx.GetAlbumById(id)
}

func (s *albumService) SearchAlbumsByTitle(title string) ([]*model.Album, error) {
	return s.ctx.GetAlbumListByTitle(title)
}

func (s *albumService) DeleteAlbum(id string) error {
	return s.ctx.DeleteAlbum(id)
}
