# epochs
Convert various epoch times to time.Time in Go.

```go
package main

import (
	"fmt"
	"time"

	"github.com/oylenshpeegul/epochs"
)

func main() {

	u := epochs.Unix(1234567890)
	fmt.Println(u.Format(time.RFC3339))

	c := epochs.Chrome(12879041490654321)
	fmt.Println(c.Format(time.RFC3339Nano))

}
```
