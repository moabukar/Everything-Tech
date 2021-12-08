# GoLang Resources and Projects

# Golang-learning

Just a repo for my personal GoLang learning and progression :)

- [Install Go here](https://go.dev/doc/install)

## Some useful GO CLI commands

```sh
- go version (prints current running Go version)

- go env (prints env vars)

- go fmt (formats code of all files in current directory)

- go build (compiles a bunch of go source code files)

- go run . (compiles and executes your code)

- go install (compiles the program & "installs" a package)

- go mod tidy (remove unused dependencies)

- go get (downloads raw source code of a 3rd party package)

- go mod init example/test (creates a new module, initialising go.mod)

- go test (runs any tests associated with the current project)

- go list (lists all currnet module dependencies)

- go help (list of all commands)
```

## managing & upgrading dependencies (example) 

```sh
- go get goloang.org/x/text (example)
- go list -m all (list all dependencies)
- go get rsc.io/sampler
- go list -m -version rsc.io/sampler (search for versions)
- go get rsc.io/sampler@v1.3.1
```


## Some other programming knowledge and stuff (just some random notes)

- dynamically typed: a language where the developer essentially doesn't care what value we assign a variable

examples: JS, Ruby & Python

- statically typed: variable values must be allocated a data type

examples: C++, Java & Go

## Links & docs

- [golang.org](golang.org)

- [godoc.org](godoc.org)

- [docs for Go packages](golang.org/pkg)

## Maintainers of this repo

- Name 1
- Name 2
