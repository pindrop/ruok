# ruok

`ruok` is a simple HTTP server that responds `200 OK` to requests to any path.

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

## Background

`ruok` was written as a utility to aid in testing the reachability of services behind an API gateway. It is intended to be a low-cost health check separate from other services. In other words, it answers when we ask "r u ok".

## (Near) Future Improvements

Our desire is to keep this service very simple, but there are some things we think could be valuable additions.

- **Configurability** - Let the user set the port and anything else that might make it easier to run.
- **Additional endpoints**
  - **Status code** - for example, `GET /status/403` returns a `403`.
  - **Headers** - Respond with request headers.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)
