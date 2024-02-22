# cython Go integration experiment

Run with:

```bash
go generate . && go run .
```

This experiment shows how to write a C function in cython that calls something requiring the python runtime to be working (in this case `print()`)
and make it callable from Go via `cgo`
