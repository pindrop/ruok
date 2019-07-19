# ruok

ruok is a simple HTTP server that responds `200 OK` to requests to any path. It is meant to be used as part of reachability tests for web APIs.

Build it:

```bash
go test -v
go build
```

Run it:

```bash
./ruok
```

Now send some requests!

```bash
curl localhost:8081
# {"OK": true}
```

Looks like everything is going to be `200 OK`.
