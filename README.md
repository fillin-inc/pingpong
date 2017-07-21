# PingPong handler

`pingpong` is a package for Go. This package provides handlers(`pingpong.Handler`, `pingpong.HandlerJSON`) for API continuity confirmation.

## Example

``` golang
package main

import (
	"net/http"

	"github.com/fillin-inc/pingpong"
)

func main() {
	http.HandleFunc("/ping", pingpong.Handler) // return text/plain
	http.HandleFunc("/ping-json", pingpong.HandlerJSON) // return application/json
	http.ListenAndServe(":8080", nil)
}
```

## License

MIT License. See [LICENSE](https://github.com/fillin-inc/pingpong/blob/master/LICENSE) file.
