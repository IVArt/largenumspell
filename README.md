# largenumspell

A small dependency free library to spell out numeric strings. It supports large numbers up to a centillion.

# Usage
```go
package main

import (
    "fmt"
    "log"

    "github.com/IVArt/largenumspell"
)

func main() {
    temp, err := largenumspell.Spellout("101")
    if err == nil {
        fmt.Printf("%#v\n", temp) // []string{"one", "hundred", "one"}
    } else {
        log.Fatal(err)
    }
    temp, err = largenumspell.Spellout("1.02e-10")
    if err == nil {
        fmt.Printf("%#v\n", temp) // []string{"one", "and", "two", "hundredth", "exponent", "minus", "ten"}
    } else {
        log.Fatal(err)
    }
    cent := "-123000"
    for i := 0; i < 30; i++ {
        cent += "0000000000"
    }
    temp, err = largenumspell.Spellout(cent)
    if err == nil {
        fmt.Printf("%#v\n", temp) // []string{"minus", "one", "hundred", "twenty-three", "centillion"}
    } else {
        log.Fatal(err)
    }
    small := "0.000"
    for i := 0; i < 30; i++ {
        small += "0000000000"
    }
    small += "45"
    temp, err = largenumspell.Spellout(small)
    if err == nil {
        fmt.Printf("%#v\n", temp) // []string{"zero", "and", "fourty-five", "hundred-centillionth"}
    } else {
        log.Fatal(err)
    }
}
```
