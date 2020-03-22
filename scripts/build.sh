set GOARCH=amd64
set GOOS=windows
go build -ldflags "-s -w" -o ../dist/ytdlgo-windows-amd64.exe ./cmd/main

set GOARCH=386
set GOOS=windows
go build -ldflags "-s -w" -o ../dist/ytdlgo-windows-386.exe ./cmd/main

set GOARCH=amd64
set GOOS=linux
go build -ldflags "-s -w" -o ../dist/ytdlgo-linux-amd64 ./cmd/main

set GOARCH=386
set GOOS=linux
go build -ldflags "-s -w" -o ../dist/ytdlgo-linux-386 ./cmd/main

set GOARCH=arm64
set GOOS=linux
go build -ldflags "-s -w" -o ../dist/ytdlgo-linux-arm64 ./cmd/main

set GOARCH=arm
set GOOS=linux
go build -ldflags "-s -w" -o ../dist/ytdlgo-linux-arm ./cmd/main

set GOARCH=arm64
set GOOS=linux
go build -ldflags "-s -w" -o ../dist/ytdlgo-linux-arm64 ./cmd/main

set GOARCH=arm
set GOOS=linux
go build -ldflags "-s -w" -o ../dist/ytdlgo-linux-arm ./cmd/main

set GOARCH=amd64
set GOOS=darwin
go build -ldflags "-s -w" -o ../dist/ytdlgo-darwin-amd64 ./cmd/main

set GOARCH=386
set GOOS=darwin
go build -ldflags "-s -w" -o ../dist/ytdlgo-darwin-386 ./cmd/main
