step set up swaggo :
go mod init swaggo
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
kalo masih gabisa run swag, add path
export PATH=$(go env GOPATH)/bin:$PATH

go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template
go get -u github.com/gorilla/mux


generate swagger/docs : swag init -g main.go
cek swagger docs : http://localhost:8080/swagger/index.html/
