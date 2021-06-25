package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	bs, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	blocksize := int(bs)
	totalsize := 10 * 1024 * 1024 * 1024
	blocknum := totalsize / blocksize

	data := make([]byte, blocksize)
	for i := 0; i < blocksize; i++ {
		data[i] = byte(i)
	}

	l, err := net.Listen("tcp4", ":22334")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < blocknum; i++ {
		n, err := conn.Write(data)
		if n != blocksize || err != nil {
			log.Fatal(err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}
