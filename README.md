# epochs
Convert various epoch times to `time.Time` times in Go.

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

**Update:** Now there are functions in the other direction too! For example, running this

```go
	fmt.Println(epochs.ToUnix(u))

	fmt.Println(epochs.ToChrome(c))
```

gives

```
1234567890
12879041490654321
```

## Contributors

[@noppers](https://github.com/noppers) originally worked out how to do the Google calendar calculation.

## See also

This project was originally done in [Perl](https://github.com/oylenshpeegul/Epochs-perl). See [the Epochs page](http://oylenshpeegul.github.io/Epochs-perl/) for motivation.

There are also a versions in
- [Elixir](https://github.com/oylenshpeegul/Epochs-elixir)
- [PowerShell](https://github.com/oylenshpeegul/Epochs-powershell)

