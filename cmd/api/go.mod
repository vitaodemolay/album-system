module github.com/vitaodemolay/album-system/cmd/api

go 1.23.3

require github.com/gorilla/mux v1.8.1

require github.com/vitaodemolay/album-system/cmd/api/controller v1.0.0

replace github.com/vitaodemolay/album-system/cmd/api/controller => ./controllers
