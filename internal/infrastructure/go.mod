module github.com/vitaodemolay/album-system/internal/infrastructure

go 1.23.3

require (
	github.com/confluentinc/confluent-kafka-go v1.9.2
	github.com/microsoft/go-mssqldb v1.8.0
	github.com/stretchr/testify v1.10.0
	github.com/vitaodemolay/album-system/internal/model v1.0.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.35.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/vitaodemolay/album-system/internal/model => ../model
