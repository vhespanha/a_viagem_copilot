PREFIX != go env GOROOT

.PHONY: build serve clean lint

build: clean
	@GOOS=js GOARCH=wasm go build -o build/hello.wasm
	@cp $(PREFIX)/lib/wasm/wasm_exec.js build/
	@cp -r static/* build/

serve: build
	@python3 -m http.server -d build

format:
	@go fmt ./...

clean:
	@rm -rf build/