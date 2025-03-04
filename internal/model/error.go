package model

import "errors"

var (
	ErrNoRow                  = errors.New("no rows in result set")
	ErrAlbumIdNotInformed     = errors.New("id field is required")
	ErrAlbumTitleNotInformed  = errors.New("title field is required")
	ErrAlbumArtistNotInformed = errors.New("artist field is required")
	ErrAlbumPriceIsInvalid    = errors.New("price field must be greater than zero")
)

type HTTPError struct {
	Code    int    `json:"code"  examples:"400"`
	Message string `json:"message" examples:"status bad request"`
}
