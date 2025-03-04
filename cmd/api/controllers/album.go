package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitaodemolay/album-system/internal/model"
)

func (ctrl *Controller) GetAlbums(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	var albums []*model.Album
	var err error

	if title == "" {
		albums, err = ctrl.albumService.GetAllAlbums()
		if err != nil {
			log.Output(0, err.Error())
			http.Error(w, "Fail on get albums", http.StatusInternalServerError)
			return
		}
	} else {
		albums, err = ctrl.albumService.SearchAlbumsByTitle(title)
		if err != nil {
			log.Output(0, err.Error())
			http.Error(w, "Fail on search albums by title "+title, http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func (ctrl *Controller) GetAlbumById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		http.Error(w, "Album id not provided", http.StatusBadRequest)
		return
	}

	album, err := ctrl.albumService.GetAlbumById(params["id"])
	if err != nil {
		log.Output(0, err.Error())
		http.Error(w, "Fail on get album with id "+params["id"], http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(album)
}

func (ctrl *Controller) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		http.Error(w, "Album id not provided", http.StatusBadRequest)
		return
	}

	err := ctrl.albumService.DeleteAlbum(params["id"])
	if err != nil {
		log.Output(0, err.Error())
		http.Error(w, "Fail on delete album with id "+params["id"], http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("deleted"))
}

func (ctrl *Controller) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album *model.Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		log.Output(0, err.Error())
		http.Error(w, "Fail on decode album", http.StatusBadRequest)
		return
	} else if err = album.Validation(); err != nil {
		log.Output(0, err.Error())
		http.Error(w, "Fail on validate album", http.StatusBadRequest)
		return
	}

	id, err := ctrl.albumService.AddAlbum(album)
	if err != nil {
		log.Output(0, err.Error())
		http.Error(w, "Fail on create album", http.StatusInternalServerError)
		return
	} else if id <= 0 {
		log.Output(0, "Album not persisted")
		http.Error(w, "Fail on persist album", http.StatusInternalServerError)
		return
	}
	album.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(album)
}
