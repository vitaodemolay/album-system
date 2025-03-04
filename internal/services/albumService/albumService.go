package services

import (
	"github.com/vitaodemolay/album-system/internal/model"
	services "github.com/vitaodemolay/album-system/internal/services/dbcontext"
)

type AlbumService struct {
	ctx services.ISqlDbContext
}

func NewAlbumService(conString string) (*AlbumService, error) {
	ctx, err := services.NewSqlDbContext(conString)

	if err != nil {
		return nil, err
	}

	return &AlbumService{ctx: ctx}, nil
}

func (s *AlbumService) AddAlbum(album *model.Album) error {
	return s.ctx.AddAlbum(album)
}

func (s *AlbumService) GetAllAlbums() ([]*model.Album, error) {
	return s.ctx.GetAlbumListByTitle("")
}

func (s *AlbumService) GetAlbumById(id string) (*model.Album, error) {
	return s.ctx.GetAlbumById(id)
}

func (s *AlbumService) SearchAlbumsByTitle(title string) ([]*model.Album, error) {
	return s.ctx.GetAlbumListByTitle(title)
}

func (s *AlbumService) DeleteAlbum(id int) error {
	return s.ctx.DeleteAlbum(id)
}
