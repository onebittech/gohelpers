# gohelpers
Helper functions for golang projects

## Mssql Helpers

```go
// MssqlDecimalToFloat64 Converts a Mssql decimal base 64 binary representation to float64
MssqlDecimalToFloat64(val string, mantissa uint8) (float64, error)
```

## Parallel Helpers
```go
// ParallelIterator Iterates a slice in parallel
ParallelIterator(length int, exec func(i int) error) error
```
