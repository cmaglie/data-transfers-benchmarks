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
	ts, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	blocksize := int(bs)
	totalsize := int(ts)
	blocknum := totalsize / blocksize

	data := make([]byte, blocksize)
	for i := 0; i < blocksize; i++ {
		data[i] = byte(i)
	}

	l, err := net.Listen("tcp4", "127.0.0.1:22334")
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
