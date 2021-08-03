.SILENT:

goversion:
	go version

build: goversion
	go build -o ./bin/akina ./cmd/akina/main.go

run: build
	./bin/akina
