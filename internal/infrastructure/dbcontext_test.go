package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitaodemolay/album-system/internal/model"
)

const (
	connectionString = "sqlserver://sa:PassW0rd@localhost:5433?database=SLQ_ALBUMSYSTEM_DB&connection+timeout=30"
)

func TestAddAlbum(t *testing.T) {
	// Arrange
	ctx, err := NewSqlDbContext(connectionString)
	assert.NoError(t, err)
	var id int

	// Act
	id, err = ctx.AddAlbum(&model.Album{ID: 1, Title: "Test Album", Artist: "Test Artist", Price: 10.0})

	// Assert
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, id, 0)
}

func TestGetAlbumById(t *testing.T) {
	// Arrange
	ctx, err := NewSqlDbContext(connectionString)
	assert.NoError(t, err)

	// Act
	album, err := ctx.GetAlbumById("1")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int(1), album.ID)
	assert.Equal(t, "Test Album", album.Title)
	assert.Equal(t, "Test Artist", album.Artist)
	assert.Equal(t, 10.0, album.Price)
}

func TestGetAlbumListByTitle(t *testing.T) {
	// Arrange
	ctx, err := NewSqlDbContext(connectionString)
	assert.NoError(t, err)

	// Act
	albums, err := ctx.GetAlbumListByTitle("Test")

	// Assert
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(albums), 1)
	assert.Equal(t, int(1), albums[0].ID)
	assert.Equal(t, 10.0, albums[0].Price)
}

func TestDeleteAlbum(t *testing.T) {
	// Arrange
	ctx, err := NewSqlDbContext(connectionString)
	assert.NoError(t, err)

	// Act
	err = ctx.DeleteAlbum("3")

	// Assert
	assert.NoError(t, err)
}
