# Central bank of the Russian Federation currency rates client [CBRF](http://www.cbr.ru/scripts/Root.asp?PrtId=SXML)

## Usage example

```golang
package main

import (
    cbr "github.com/matperez/go-cbr-client"
    "time"
    "fmt"
)

func main() {
    client := cbr.NewClient()
    rate, err := client.GetRate("USD", time.Now())
    if err != nil {
        panic(err)
    }
    fmt.Prinln(rate)
}
```
