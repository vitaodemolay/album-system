package infrastructure

import (
	"context"
	"database/sql"

	mssql "github.com/microsoft/go-mssqldb"
	"github.com/vitaodemolay/album-system/internal/model"
)

type ISqlDbContext interface {
	AddAlbum(album *model.Album) (int, error)
	GetAlbumById(id string) (*model.Album, error)
	GetAlbumListByTitle(title string) ([]*model.Album, error)
	DeleteAlbum(id string) error
}

type sqlDbContext struct {
	Db *sql.DB
}

func NewSqlDbContext(conString string) (*sqlDbContext, error) {
	db, err := sql.Open("sqlserver", conString)
	if err != nil {
		return nil, err
	}

	return &sqlDbContext{Db: db}, nil
}

func (ctx *sqlDbContext) AddAlbum(album *model.Album) (int, error) {
	scriptSql := "INSERT INTO TbAlbum (Title, Artist, Price) OUTPUT INSERTED.ID VALUES (@Title, @Artist, @Price)"
	var id int
	err := ctx.Db.QueryRowContext(context.Background(), scriptSql, sql.Named("Title", mssql.VarChar(album.Title)), sql.Named("Artist", mssql.VarChar(album.Artist)), sql.Named("Price", album.Price)).Scan(&id)
	return id, err
}

func (ctx *sqlDbContext) GetAlbumById(id string) (*model.Album, error) {
	scriptSql := "SELECT ID, Title, Artist, Price FROM TbAlbum WHERE ID = @ID"
	row := ctx.Db.QueryRow(scriptSql, sql.Named("ID", id))

	album := new(model.Album)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return album, nil
	}
}

func (ctx *sqlDbContext) GetAlbumListByTitle(title string) ([]*model.Album, error) {
	scriptSql := "SELECT ID, Title, Artist, Price FROM TbAlbum WHERE Title LIKE @Title"
	rows, err := ctx.Db.Query(scriptSql, sql.Named("Title", mssql.VarChar("%"+title+"%")))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := make([]*model.Album, 0)
	for rows.Next() {
		album := new(model.Album)
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (ctx *sqlDbContext) DeleteAlbum(id string) error {
	scriptSql := "DELETE FROM TbAlbum WHERE ID = @ID"

	_, err := ctx.Db.Exec(scriptSql, sql.Named("ID", id))
	return err
}
