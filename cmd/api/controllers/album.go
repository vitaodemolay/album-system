package controller

import "net/http"

func (ctrl *Controller) GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of albums"))
}

func (ctrl *Controller) GetAlbumById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get Album by id"))
}

func (ctrl *Controller) SearchAlbumsByTitle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Search Albums by title"))
}

func (ctrl *Controller) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Album"))
}

func (ctrl *Controller) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Album"))
}
