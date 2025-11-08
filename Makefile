PREFIX != go env GOROOT

.PHONY: build serve format check clean devtools

build: clean check
	@GOOS=js GOARCH=wasm go build -o build/hello.wasm
	@cp $(PREFIX)/lib/wasm/wasm_exec.js build/
	@cp -r static/* build/

serve: build
	@python3 -m http.server -d build

format:
	@goimports -w .

check: format
	@GOOS=js GOARCH=wasm go vet ./...
	@GOOS=js GOARCH=wasm staticcheck -checks "all" ./...

clean:
	@rm -rf build/

devtools:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install golang.org/x/tools/cmd/goimports@latest
