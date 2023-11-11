build:
	@go build -ldflags "-s -w -H windowsgui" -o ./bin/game.exe ./cmd/main.go 

build-static:
	@go build -tags static -ldflags "-s -w -H windowsgui" -o ./bin/game-static.exe ./cmd/main.go  

run: build
	@./bin/game.exe

test:
	@go test -v ./...

