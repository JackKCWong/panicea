# panicea

An anti-idiom error handling for Go 1 : just panic!

## usage

```go get github.com/JackKCWong/panicea```

```go
import (
    . "github.com/JackKCWong/panicea"
)

func main() {
    err := Try(func() {
        panic(fmt.Errorf("just panic with an error"))
    })

    if err != nil {
        log.Fatal(err)
    }

    // or if you need return value
    r, err := Trap(func() int {
        Check(somethingMayErr())            // panic if err
        r = Must(somethingMayReturnIntOrErr()) // panic if err

        return r
    })

    if err != nil {
        log.Fatal(err)
    }
}
```
