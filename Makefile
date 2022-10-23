serve: compile
	go run server.go

compile: main.go
	GOARCH=wasm GOOS=js go build -o test.wasm main.go
