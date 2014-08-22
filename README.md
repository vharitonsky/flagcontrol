flagcontrol
===========

Web interface to see and control flags in your app 

Usage
======

```go

package main

import (
	"flag"
	"fmt"
	"github.com/vharitonsky/flagcontrol"
	"log"
	"net/http"
)

var (
	fl1 = flag.String("fl1", "v1", "Testflag")
)

func main() {
	http.HandleFunc("/flags", flagcontrol.Server)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


```

Now you can navigate to localhost/flags and view/set your flags live.
