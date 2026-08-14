// This file lives outside the go/ module for parity with the other
// language examples. To run it, either copy it into go/cmd/example/main.go,
// or temporarily add it to the module:
//
//   cp examples/go_example.go go/cmd/example_main.go && \
//   cd go && go run ./cmd/example_main.go; rm cmd/example_main.go
package main

import (
	"fmt"

	"github.com/verhas/babbler/identifier"
)

func main() {
	// Typical usage: give a friendly display name to each row in an
	// auto-increment sequence (e.g. a database primary key).
	for userID := 0; userID < 5; userID++ {
		id, err := identifier.NumberToID(userID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("user #%d -> %s\n", userID, id)
	}
}
