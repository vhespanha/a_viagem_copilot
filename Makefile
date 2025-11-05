PREFIX != go env GOROOT

build: clean
	@GOOS=js GOARCH=wasm go build -o build/hello.wasm
	@cp $(PREFIX)/lib/wasm/wasm_exec.js build/
	@cp -r static/* build/

serve: build
	@python3 -m http.server -d build

clean:
	@rm -rf build/
