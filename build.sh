GOOS=linux
GOARCH=amd64
go build -o ./binary/jan-go_${GOOS}_${GOARCH}
GOOS=windows
GOARCH=amd64
go build -o ./binary/jan-go_${GOOS}_${GOARCH}.exe
