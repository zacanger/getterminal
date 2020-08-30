# getterminal

[![Support with PayPal](https://img.shields.io/badge/paypal-donate-yellow.png)](https://paypal.me/zacanger) [![Patreon](https://img.shields.io/badge/patreon-donate-yellow.svg)](https://www.patreon.com/zacanger) [![ko-fi](https://img.shields.io/badge/donate-KoFi-yellow.svg)](https://ko-fi.com/U7U2110VB)

--------

Get the best terminal the user has installed. Known to work on Linux, does not
work on Windows or Mac yet. Go port of my Node package
[get-term](https://npm.im/get-term).

## Installation

`go get github.com/zacanger/getterminal`

## Usage

```go
package main

import (
  "fmt"
  "github.com/zacanger/get-terminal"
)

func main() {
  bestTerminal := getterminal.GetTerminal()
  fmt.Println("best terminal to use is", bestTerminal)
}
```

[LICENSE](./LICENSE.md)
