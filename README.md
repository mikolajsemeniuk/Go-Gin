# Supreme-Go
get started
```sh
brew install go
go version
echo $GOROOT
which go # /usr/local/bin/go
~/.zshrc
export GOROOT=/usr/local/Cellar/go/1.17.6/libexec
export PATH=$PATH:$GOROOT/bin
source ~/.zshrc


go mod init github.com/mikolajsemeniuk/Supreme-Go
go run main.go

# generate new swagger
swag init
# http://localhost:8080/swagger/index.html#/
# http://localhost:8080/swagger/doc.json

go test ./...
go test ./controllers/account 
```