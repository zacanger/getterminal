# getterminal

--------

Get the best terminal the user has installed. Go port of my Node package
[get-term](https://npm.im/get-term).

## Installation

`go get github.com/zacanger/getterminal`

## Usage

```go
package main

import (
  "fmt"
  "os/exec"

  "github.com/zacanger/get-terminal"
)

func main() {
  bestTerminal := getterminal.GetTerminal()
  fmt.Println("best terminal to use is", bestTerminal)
  exec.Cmd(bestTerminal).Run()
}
```

## Contributing

Contributions are welcome! See [CONTRIBUTING](/.github/CONTRIBUTING.md)

[LICENSE](./LICENSE.md)
