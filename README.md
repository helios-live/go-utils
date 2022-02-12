# go-utils

Some simple utilities that make life simpler when developing with go

### utils.Time
```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ideatocode/go-utils"
)

func main() {
	// Create a new Time
	tm := utils.Time{Time: time.Now()}

	// marshal it
	bytes, _ := json.Marshal(tm)
	fmt.Println(string(bytes))

	var tm2 utils.Time

	// unmarhal it
	json.Unmarshal(bytes, &tm2)
	fmt.Println(tm2.String())
}
```
