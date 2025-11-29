# golang-starterkit

## Running example modules (chapters)

This repository contains small example modules (each is a directory containing a `main.go`) you can run while following a book or tutorial. Use the provided `Makefile` to discover, run and build these modules easily.

Common commands from the repository root:

- List detected modules (directories containing a `main.go`):

```bash
make list
```

- Run a single module (replace `DIR` with the module directory):

```bash
make run DIR=hello-world
```

- Run every module found (runs them sequentially, stops on first failure):

```bash
make run-all
```

- Build every module into `./bin` (binary names use `-` for `/`):

```bash
make build-all
```

- Run all Go tests in the repository:

```bash
make test
```

Notes:
- The Makefile detects modules by searching for `main.go`. When running/building it `cd`s into the module directory before invoking `go run`/`go build`, so you don't need a `go.mod` at the repository root.
- If you prefer a `cmd/` layout for chapters (e.g. `cmd/chapter01/main.go`) the Makefile will detect them as separate modules as well.

If you'd like, I can also add a `cmd/` scaffold command to the Makefile to create new chapter directories with a starter `main.go`.