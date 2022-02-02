# Supreme-Go
get started
```sh
brew install go
go version
echo $GOROOT
which go # /usr/local/bin/go
# ~/.zshrc
# export GOROOT=/usr/local/Cellar/go/1.17.6/libexec
# export PATH=$PATH:$GOROOT/bin
export GOPATH=/Users/mikolajsemeniuk/go
export PATH=$GOPATH/bin:$PATH
source ~/.zshrc

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

go mod init github.com/mikolajsemeniuk/Supreme-Go
go run main.go

# generate new swagger
swag init
# http://localhost:3000/swagger/index.html#/
# https://localhost:3000/swagger/index.html#/
# http://localhost:8080/swagger/doc.json

# live reload
go get -u github.com/cosmtrek/air
air

go test ./...
go test ./controllers/account 
```