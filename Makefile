# Example Makefile.
.PHONY: build clean default edit init dis run

default: run

# Build the system.
build: init www/wasm/site.wasm bin/server

# Build the webassembly module using the go compiler.
www/wasm/site.wasm: src/local/site.go
	GOARCH=wasm GOOS=js GOPATH=$$(pwd) go build -o $@ $<

# Build the http server.
#
# Run it like this (or "make run"):
#    $ ./bin/server -dir www
#
# It can then be accessed from the browser as localhost:5555.
bin/server: src/local/server.go | bin
	GOPATH=$$(pwd) go build -buildmode=exe -o $@ $<

bin:
	@mkdir $@

# Build and run the server.
run: build
	@echo -e "\033[1mINFO: Navigate to localhost:5555 for the demo.\033[0m"
	@bin/server -dir www

# Disassemble.
dis: build
	bin/wasm-dump -d www/wasm/site.wasm

# Clean up.
clean:
	find . -type f -name '*~' -delete
	rm -rf wasm .init www/wasm/site.wasm bin $$(ls -1d src/* | grep -v local)

# Run the atom editor on my system.
edit: init
	GOPATH=$$(pwd) /opt/atom/latest/Atom.app/Contents/MacOS/Atom

# Initialize some stuff locally.
init: .init
	GOPATH=$$(pwd) go version

# https://blog.gopheracademy.com/advent-2017/go-wasm/
.init:
	@rm -f .init
	GOPATH=$$(pwd) go get -u github.com/derekparker/delve/cmd/dlv
	GOPATH=$$(pwd) go get github.com/go-interpreter/wagon/cmd/wasm-dump
	@touch .init
