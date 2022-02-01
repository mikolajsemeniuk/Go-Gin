# Supreme-Go
get started
```sh
go mod init github.com/mikolajsemeniuk/Supreme-Go
go run main.go

# generate new swagger
swag init
# http://localhost:8080/swagger/index.html#/
# http://localhost:8080/swagger/doc.json

# mocks
go generate ./...
go tests
```