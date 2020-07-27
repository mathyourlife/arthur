# arthur

This package is intended to provide some tools to aide
in working with mathematics content.

## Examples

### HTTP Server

Sample server that currently responds with a LaTeX formatted
fraction.

```
go run github.com/mathyourlife/arthur/examples/server
```

```
curl localhost:8080
\frac{1}{18}
```
