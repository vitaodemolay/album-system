package model

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (album Album) Validation() error {
	switch {
	case album.ID <= 0:
		return ErrAlbumIdNotInformed
	case len(album.Title) == 0:
		return ErrAlbumTitleNotInformed
	case len(album.Artist) == 0:
		return ErrAlbumArtistNotInformed
	case album.Price <= 0:
		return ErrAlbumPriceIsInvalid
	default:
		return nil
	}
}
