module github.com/vitaodemolay/album-system/cmd/api/controller

go 1.23.3

require (
	github.com/gorilla/mux v1.8.1
	github.com/vitaodemolay/album-system/internal/infrastructure v1.0.0
	github.com/vitaodemolay/album-system/internal/model v1.0.0
	github.com/vitaodemolay/album-system/internal/services v1.0.0
)

require (
	github.com/confluentinc/confluent-kafka-go v1.9.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/microsoft/go-mssqldb v1.8.0 // indirect
	golang.org/x/crypto v0.35.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

replace github.com/vitaodemolay/album-system/internal/infrastructure => ../../../internal/infrastructure

replace github.com/vitaodemolay/album-system/internal/services => ../../../internal/services

replace github.com/vitaodemolay/album-system/internal/model => ../../../internal/model
