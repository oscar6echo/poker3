# Poker

> WIP

```bash
# test
# from /poker
go test -v

# benchmark
# from /poker
go test -run=XXX -bench=.

# main ie other test
# from top repo
go run main.go

# compile - tinygo
tinygo build -o ./html/main_js.wasm -target wasm ./main.go
tinygo build -o ./html/main_js.wasm -target wasm -no-debug ./main.go

# serve
go run serve/serve.go
# open http://localhost:8083
```
