# strftime

Go date time format compatible with Ruby's [Time#strftime](http://ruby-doc.org/core-2.2.2/Time.html#method-i-strftime)

## Example

```go
package main

import (
  "fmt"
  "time"

  "github.com/acstech/liquid/strftime"
)

func main() {
  t := time.Now()
  s := strftime.Strftime(&t, "%Y-%m-%d %H:%M:%S")
  fmt.Println(s)
}
```
