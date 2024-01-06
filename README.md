# vnd

Template for REST API server made with Golang. This template is built on top of [`echo`](https://echo.labstack.com/).

Features:
- Graceful stop
- Close connections before stop
- Response template error/success/success with pagination

## Installation

To install `vnd` package, you need to install Go.

1. You first need [Go](https://golang.org/) installed then you can use the below Go command to install `vnd`.

```sh
go get -u github.com/thienhaole92/vnd
```

2. Import it in your code:

```go
import "github.com/thienhaole92/vnd"
```

## Quick start

### Starting HTTP server

The example how to use `vnd` with to start REST api in [example](./example) folder.
