# search

![Go](https://github.com/qba73/search/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/search)](https://goreportcard.com/report/github.com/qba73/search)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/search)
![GitHub](https://img.shields.io/github/license/qba73/search)

```go
import github.com/qba73/search
```

# What is search?

Search is a small utility package for searching input for a given sub-string.

# How to use it?

The search uses standard input and output as default settings (`stdin`, `stdout`).

```go
Search("hello")
```


Use functionl options to change default settings and use different sources and destinations.

```go
s := search.New(
    WithInput(io.Reader),
    WithOutput(io.Writer),
)

s.Search("hello")
```



