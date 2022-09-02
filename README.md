# monkey-go
Building an interpreter and compiler by following Thorsten Ball's books

------------------------------

**Run the repl**

```bash
cd interpreter
go run main.go
```

**Build the web repl**

```bash
cd interpreter/web-repl
GOOS=js GOARCH=wasm go build -o ./public/monkey-repl.wasm 
```

**Serve the web repl on port 9000**

```bash
cd interpreter/web-repl/public
python3 -m http.server 9000 --bind 0.0.0.0
```

