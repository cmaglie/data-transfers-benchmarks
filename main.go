package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"time"

	"github.com/arduino/go-paths-helper"
)

func main() {
	bs := []int{
		1024,
		2 * 1024,
		4 * 1024,
		8 * 1024,
		16 * 1024,
		32 * 1024,
		64 * 1024,
		128 * 1024,
		256 * 1024,
		512 * 1024,
		1024 * 1024}
	for _, s := range bs {
		TestStdio(s)
	}
	for _, s := range bs {
		TestTCP(s)
	}
}

func TestStdio(blocksize int) {
	tmp, err := paths.MkTempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.RemoveAll()
	srv := tmp.Join("stdio")
	cmd := exec.Command("go", "build", "-o", srv.String(), "./stdio-server")
	if err := cmd.Run(); err != nil {
		log.Fatal("Run: ", err)
	}

	s := exec.Command(srv.String(), fmt.Sprintf("%d", blocksize))
	out, err := s.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	w, err := io.Copy(ioutil.Discard, out)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("stdio BS=%7d: Read %d bytes in %7.3f sec: %7.3f MB/sec\n", blocksize, w, elapsed.Seconds(), float64(w)/elapsed.Seconds()/1024/1024)

	if err := out.Close(); err != nil {
		log.Fatal(err)
	}
	if err := s.Wait(); err != nil {
		log.Fatal(err)
	}
}

func TestTCP(blocksize int) {
	tmp, err := paths.MkTempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.RemoveAll()
	srv := tmp.Join("tcp")
	cmd := exec.Command("go", "build", "-o", srv.String(), "./tcp-server")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	s := exec.Command(srv.String(), fmt.Sprintf("%d", blocksize))
	go s.Run()

	time.Sleep(500 * time.Millisecond)
	conn, err := net.Dial("tcp4", "127.0.0.1:22334")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	w, err := io.Copy(ioutil.Discard, conn)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("tcpip BS=%7d: Read %d bytes in %7.3f sec: %7.3f MB/sec\n", blocksize, w, elapsed.Seconds(), float64(w)/elapsed.Seconds()/1024/1024)

	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}
