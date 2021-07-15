# Poker

## Overview

This repo contains:

- [poker](./poker): A **Go package** to evaluate Poker hands - Texas Hold'em
- [keygen-go](./keygen-go): A **Go script** to search for keys used in package **poker**
- [keygen-c](./keygen-c): A **C script** to search for keys used in package poker - faster than **keygen-go**
- [serve](./serve): A **Go script** to serve folder [html](./html)
- [html](./html): A minimal **web page** using a **WASM module** compiled from script [main.go](./main.go) using package **poker**
- [wip](./wip): A sample **Go script** using package **poker**

Ref:

- repo [oscar6echo/tinygo-wasm-demo-2](https://github.com/oscar6echo/tinygo-wasm-demo-2)
- [Go package syscall/js doc](https://pkg.go.dev/syscall/js)
- ObservableHQ notebook [Poker Hand Evaluator - WASM](https://observablehq.com/@oscar6echo/poker-hand-evaluator-wasm)

## Poker Go package

This [poker](./poker) Go package uses an algorithm introduced by [SpecialK](https://github.com/kennethshackleton/SKPokerEval).  
The underlying idea is to precalculate everything so that evaluating a hand is at most a couple of table lookups. Obviously this is very fast.

Cards are identified by their faces (2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A) only for _non flush hands_. So any combination of (7 among 13) should yield a different index. This could be done by mapping each face to one of the first 13 prime numbers, by example. But then the max index would be 2x3x5x7x11x..x41 ~ 3e14 which is way to large to fit in memory. So another mechanism must be found. SpecialK replaces multiplication by addition and looks for keys, such that no two distinct combinations (with max 4 identical keys) give the same sum, by brute force. Fortunately these keys (the "face keys") exist and can be found in a few hours. See [keygen-c](./keygen-c).

But first to determine whether a hand is a flush, apply the same mechanism to a much smaller problem: Each suit (spades, hearts, diamonds, clubs) is mapped to a key so that no two distinct combinations of 4 suits (with max 4 identical keys) add up to the same index. These keys (the "suit keys") are small and found instantly.

If the hand is not a flush, the keys described above are used. If it is, then again apply the same mechanism to a small problem: Each face is mapped to a key so that no two distinct combinations of 7 keys (all distinct) add up to the same index. These keys (the "flush keys") are small and found in a few seconds.

Besides it turns out these keys are small enough so that the face keys and the suit keys together fit into 32-bit integers.  
This is a suprising coincidence !

### Test

```bash
# test
# from /poker
go test -v

# benchmark
# from /poker
go test -run=XXX -bench=.

# sample script
# from /wip
go run sample.go
```

The package is quite fast as it goes through all (7 among 52) = 133.8m cases in 4s seconds.

```bash
go test -v
=== RUN   TestGetRankFive
--- PASS: TestGetRankFive (0.02s)
=== RUN   TestGetRankSeven
--- PASS: TestGetRankSeven (0.02s)
=== RUN   TestGetRank
--- PASS: TestGetRank (0.02s)
=== RUN   TestBuildFiveHandStats
--- PASS: TestBuildFiveHandStats (0.11s)
=== RUN   TestBuildSevenHandStats
--- PASS: TestBuildSevenHandStats (3.96s)
PASS
ok      github.com/oscar6echo/poker3/poker      4.132s
```

### WASM

To try and make this package available for the web, compile it with [TinyGo](https://tinygo.org/).

```bash
# compile - tinygo
tinygo build -o ./html/main_js.wasm -target wasm ./main.go

# compile - w/o debug
tinygo build -o ./html/main_js.wasm -target wasm -no-debug ./main.go
```

Then serve folder [html](./html) containing a minimal web page with the compiled WASM module.

```bash
# serve
go run serve/serve.go
# open http://localhost:8083
```

Portability is achieved as the poker package is exposed to a web page.  
But performance is **surprisingly poor**: function `TestBuildFiveHandStats` takes about 50s in the webpage vs. 0.1s in native Go !!!

_NOTE_: Script `wasm_exec.js` is used in web page [html/index.html](./html/index.html). `wasm_exec_mod.js` is its conversion to a JavaScript module, used in ObservableHQ notebook [Poker Hand Evaluator - WASM](https://observablehq.com/@oscar6echo/poker-hand-evaluator-wasm).
