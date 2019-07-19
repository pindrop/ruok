# ruok

ruok is a simple HTTP server that responds `200 OK` to requests to any path. It is meant to be used as part of reachability tests for web APIs.

## building

```bash
go test -v ./...
go build
```

## running

```bash
./ruok
```
