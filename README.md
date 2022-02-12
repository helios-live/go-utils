# go-utils

Some simple utilities that make life simpler when developing with go


### utils.PrintReader
An encapsulator of a reader that also prints whatever is read. 
Useful when turning up debug levels and want to see under the hood
```go
package main

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/ideatocode/go-utils"
)

func getReader(prefix string) io.Reader {

	data := "Some information"

	reader := strings.NewReader(data)
	if utils.DebugLevel > 9 {
		return utils.PrintReader{
			Reader: reader,
			Prefix: prefix,
		}
	}
	return reader
}

func main() {

	utils.DebugLevel = 1
	sr := getReader("Not printed")

	// does not print anything
	ioutil.ReadAll(sr)

	utils.DebugLevel = 420
	pr := getReader("Printed")

	// also prints
	ioutil.ReadAll(pr)
}
```

### utils.Debug
Useful Debugging utilities using log levels
```go
package main

import (
	"fmt"

	"github.com/ideatocode/go-utils"
)

func main() {
	utils.DebugLevel = 1

	// just like Println but only writes if utils.DebugLevel >= 1
	utils.Debugln(1, "This is a", "useful", "information")

	// debug, just like Printf 
	utils.Debugf(999, "This is a debugging message")

	// print it
	fmt.Println(tm)

}


```
### utils.RandStringBytes(63)
Generate a random string with go
```go
package main

import (
	"fmt"

	"github.com/ideatocode/go-utils"
)

func main() {
	// Create a new Time
	tm := utils.RandStringBytes(63)

	// print it
	fmt.Println(tm)

}
```

### utils.Time: 
A more human readable version of time.Time

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
