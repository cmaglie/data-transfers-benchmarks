package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	bs, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	blocksize := int(bs)
	const totalsize = 10 * 1024 * 1024 * 1024
	blocknum := totalsize / blocksize

	data := make([]byte, blocksize)
	for i := 0; i < blocksize; i++ {
		data[i] = byte(i)
	}

	for i := 0; i < blocknum; i++ {
		n, err := os.Stdout.Write(data)
		if n != blocksize || err != nil {
			log.Fatal(err)
		}
	}
}
