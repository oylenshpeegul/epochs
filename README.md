# epochs
Convert various epoch times to time.Time times in Go.

For example, running this code

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

would give

```
2009-02-13T23:31:30Z
2009-02-13T23:31:30.654321Z
```
