# randimgurimgurl

Getter random imgur image url implementation with Go



## Install

```bash
go get github.com/shinshin86/randimgurimgurl
```



## Usage

```go
package main

import (
	"fmt"
	"github.com/shinshin86/randimgurimgurl"
)

func main() {
	fmt.Println(randimgurimgurl.Getter())
  // https://i.imgur.com/XXXXX.png
}

```



## Test

```bash
go test
```

