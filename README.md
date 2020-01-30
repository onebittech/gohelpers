# gohelpers
Helper functions for golang projects

## How to use

```go
package main

import (
    "fmt"
    helpers "github.com/onebittech/gohelpers"
)

func main() {
    helpers.ParallelIterator(10, func (i int) error {
        fmt.Println(i)
        return nil
    })
}
```

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
