# gshogi

**gshogi** is a shogi library wrtten in Go.

This is heavily inspired by [python-shogi](https://github.com/gunyarakun/python-shogi).

## Usage

### `USIHandler`

```go
package main

import (
	"bufio"
	"log"
	"os"

	"github.com/mururu/gshogi"
)

func main() {
	engine := &gshogi.DefaultEngine{}
	handler := gshogi.NewUSIHandler(engine, os.Stdout)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		t := s.Text()
		if err := handler.Handle(t); err != nil {
			break
		}
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}
```

## Todo

- Support all USI commands
- Add helper functions to interact board data easily
- Add validation and error handling
