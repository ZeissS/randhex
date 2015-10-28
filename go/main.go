package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"flag"
)

var (
	size = flag.Int("bytes", 4, "How many bytes should be read. The number of characters in the output will be twice this value.")
)

func main() {
	flag.Parse()

	buffer := make([]byte, *size)

	_, err := rand.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", hex.EncodeToString(buffer))
}
