# go-ipinfo

Client of [ipinfo.io](http://ipinfo.io/) in Go language.

## Sample

```go
package main

import (
	"fmt"
	"net"
	"github.com/syohex/go-ipinfo"
)

func main() {
	info := ipinfo.IPInfo(net.ParseIP("111.222.333.444"))
	fmt.Printf("%v\n", info)
}
```
