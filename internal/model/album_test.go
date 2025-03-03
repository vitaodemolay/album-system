package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbum_Validation(t *testing.T) {
	tests := []struct {
		name    string
		album   Album
		wantErr error
	}{
		{
			name: "Valid Album",
			album: Album{
				ID:     1,
				Title:  "Test Album",
				Artist: "Test Artist",
				Price:  9.99,
			},
			wantErr: nil,
		},
		{
			name: "Missing ID",
			album: Album{
				Title:  "Test Album",
				Artist: "Test Artist",
				Price:  9.99,
			},
			wantErr: ErrAlbumIdNotInformed,
		},
		{
			name: "Missing Title",
			album: Album{
				ID:     1,
				Artist: "Test Artist",
				Price:  9.99,
			},
			wantErr: ErrAlbumTitleNotInformed,
		},
		{
			name: "Missing Artist",
			album: Album{
				ID:    1,
				Title: "Test Album",
				Price: 9.99,
			},
			wantErr: ErrAlbumArtistNotInformed,
		},
		{
			name: "Invalid Price",
			album: Album{
				ID:     1,
				Title:  "Test Album",
				Artist: "Test Artist",
				Price:  0,
			},
			wantErr: ErrAlbumPriceIsInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			album := tt.album

			// Act
			err := album.Validation()

			// Assert
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
