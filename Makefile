all:
	GOOS=js GOARCH=wasm go build -o  ./test.wasm
serve:
	python3 -m http.server
