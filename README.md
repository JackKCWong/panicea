# panicea

An anti-idiom error handling for Go 1 : just panic!

## usage

```
go get github.com/JackKCWong/panicea
```

```go
import (
    . "github.com/JackKCWong/panicea"
)

func main() {
    err := Trap(func() {
        panic(fmt.Errorf("just panic with an error"))
    })

    if err != nil {
        log.Fatal(err)
    }

    // or if you need return value
    v, err := Try(func() int {
        v := Catch(somethingMayReturnIntOrErr()).On("error: %w")     // panic if err

        return v
    })

    if err != nil {
        log.Fatal(err)
    }
}
```
