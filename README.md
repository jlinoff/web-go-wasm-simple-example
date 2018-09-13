# web-go-wasm-simple-example
Simple example that shows how to use go to create and use a wasm module on a web page

The goal of this example is to provide a simple, high level overview
of how to use go to create `wasm` modules and integrate them into a
browser.

It uses go to compile a source file named `site.go` (in
`src/local/site.go`) and install it in `www/wasm/site.wasm`. It then
runs a local server which serves the `www` site on `localhost:5555`.

The site displays a `Run` button which, when clicked, loads and runs
the `site.wasm` code. That code demonstrates the following actions from
the wasm code:

1. How to generate an alert box.
2. How to write console log output.
3. How to get an element by id.
4. How to set an element innerHTML property.
5. How to set an element style attributes.

### Requirements
1. go 1.11 or later
2. make

### Build and Run the Server
To build and run the server:

```bash
$ make
```

### Run the Example
To access the example page in your browser navigate to `localhost:5555`.

### Source Files and Layout

```
.
├── Makefile
├── README.md
├── src
│   └── local
│       ├── server.go
│       └── site.go
└── www
    ├── css
    │   └── site.css
    ├── index.html
    ├── js
    │   ├── site.js
    │   └── wasm_exec.js
    └── wasm

6 directories, 8 files
```

1. `server.go` is the source file for the simple web server.
2. `site.go` is the source file for the wasm code.
