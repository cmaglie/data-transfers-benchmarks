package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"runtime"
	"time"

	"github.com/arduino/go-paths-helper"
)

var osExecutableExt = ""

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
	ts := 10 * 1024 * 1024 * 1024
	if runtime.GOOS == "windows" {
		osExecutableExt = ".exe"
	}
	tmp, err := paths.MkTempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.RemoveAll()

	stdioSrv := tmp.Join("stdio" + osExecutableExt)
	{
		cmd := exec.Command("go", "build", "-o", stdioSrv.String(), "./stdio-server")
		if err := cmd.Run(); err != nil {
			log.Fatal("Run: ", err)
		}
	}

	tcpSrv := tmp.Join("tcp" + osExecutableExt)
	{
		cmd := exec.Command("go", "build", "-o", tcpSrv.String(), "./tcp-server")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	for _, s := range bs {
		TestStdio(stdioSrv, s, ts)
	}
	for _, s := range bs {
		TestTCP(tcpSrv, s, ts)
	}
}

func TestStdio(srv *paths.Path, blocksize int, totalsize int) {
	s := exec.Command(srv.String(), fmt.Sprintf("%d", blocksize), fmt.Sprintf("%d", totalsize))
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

func TestTCP(srv *paths.Path, blocksize int, totalsize int) {

	s := exec.Command(srv.String(), fmt.Sprintf("%d", blocksize), fmt.Sprintf("%d", totalsize))
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
