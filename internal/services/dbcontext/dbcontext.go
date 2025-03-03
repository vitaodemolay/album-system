package services

import (
	"database/sql"

	mssql "github.com/microsoft/go-mssqldb"
	"github.com/vitaodemolay/album-system/internal/model"
)

type SqlDbContext struct {
	Db *sql.DB
}

func NewSqlDbContext(conString string) (*SqlDbContext, error) {
	db, err := sql.Open("sqlserver", conString)
	if err != nil {
		return nil, err
	}

	return &SqlDbContext{Db: db}, nil
}

func (ctx *SqlDbContext) AddAlbum(album *model.Album) error {
	scriptSql := "INSERT INTO TbAlbum (Title, Artist, Price) VALUES (@Title, @Artist, @Price)"
	_, err := ctx.Db.Exec(scriptSql, sql.Named("Title", mssql.VarChar(album.Title)), sql.Named("Artist", mssql.VarChar(album.Artist)), sql.Named("Price", album.Price))
	return err
}

func (ctx *SqlDbContext) GetAlbumById(id string) (*model.Album, error) {
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

func (ctx *SqlDbContext) GetAlbumListByTitle(title string) ([]*model.Album, error) {
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

func (ctx *SqlDbContext) DeleteAlbum(id int) error {
	scriptSql := "DELETE FROM TbAlbum WHERE ID = @ID"

	_, err := ctx.Db.Exec(scriptSql, sql.Named("ID", id))
	return err
}
