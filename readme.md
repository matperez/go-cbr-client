# Получение курсов валют с сайта [ЦБРФ](http://www.cbr.ru/scripts/Root.asp?PrtId=SXML)

## Пример использования

```golang
package main

import (
    cbr "github.com/matperez/go-cbr-client"
    "time"
    "fmt"
)

func main() {
    rate, err := cbr.GetCurrencyRate("USD", time.Time)
    if err != nil {
        panic(err)
    }
    fmt.Prinln(rate)
}
```
