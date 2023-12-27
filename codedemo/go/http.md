```go
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "time"
)

func main() {
	cmd := os.Args[1]
	if cmd == "tcp" {
		go TCPtestReal()
		TCPtestEmu()
	} else {
		go UDPtestReal()
		UDPtestEmu()
	}
	// ln, err := net.Listen("tcp", ":43333")
	// if err != nil {
	// 	return
	// }

	// for {

	// 	conn, _ := ln.Accept()

	// 	go handle(conn)

	// }

}

func TCPtestReal() {
	ln, _ := net.Listen("tcp", ":11111")
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}

func TCPtestEmu() {
	ln, _ := net.Listen("tcp", ":11112")
	var router *mux.Router
	router = mux.NewRouter()
	router.HandleFunc("/", nodeHandler)
	var server *http.Server
	server = new(http.Server)
	server.Handler = router
	server.Serve(ln)
	// for {
	// 	conn, _ := ln.Accept()
	// 	go handleEmu(conn)
	// }
}

func nodeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write(([]byte)("EMU response"))
}

func UDPtestReal() {
	addr, _ := net.ResolveUDPAddr("udp", ":11111")
	ln, _ := net.ListenUDP("udp", addr)
	b := make([]byte, 1024)
	for {
		n, srcaddr, _ := ln.ReadFromUDP(b)
		fmt.Printf("Server get msg:%s\n", string(b[:n-1]))
		rspStr := fmt.Sprintf("Hello,%s. This is REAL device responsing in UDP.", (string)(b[:n-1]))
		ln.WriteToUDP(([]byte)(rspStr), srcaddr)
		fmt.Printf("Server send msg:%s\n", rspStr)
	}
}

func UDPtestEmu() {
	addr, _ := net.ResolveUDPAddr("udp", ":11112")
	ln, _ := net.ListenUDP("udp", addr)
	b := make([]byte, 1024)
	for {
		n, srcaddr, _ := ln.ReadFromUDP(b)
		fmt.Printf("Server get msg:%s\n", string(b[:n-1]))
		rspStr := fmt.Sprintf("Hello,%s. This is EMU device responsing in UDP.", (string)(b[:n-1]))
		ln.WriteToUDP(([]byte)(rspStr), srcaddr)
		fmt.Printf("Server send msg:%s\n", rspStr)
	}
}

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("Done")
	}()
	fmt.Println("start one")
	// s := "Server msg"
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
			}
			return
		}
		fmt.Printf("Server get msg:%s\n", string(buf[:n-1]))
		rspStr := fmt.Sprintf("Hello,%s. This is REAL device responsing in TCP.", (string)(buf[:n-1]))
		_, _ = conn.Write([]byte(rspStr))
		fmt.Printf("Server send msg:%s\n", rspStr)
		// time.Sleep(3 * time.Second)
	}
}

func handleEmu(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("Done")
	}()
	fmt.Println("start one")
	// s := "Server msg"
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
			}
			return
		}
		fmt.Printf("Server get msg:%s\n", string(buf[:n-1]))
		rspStr := fmt.Sprintf("Hello,%s. This is EMU device responsing in TCP.", (string)(buf[:n-1]))
		_, _ = conn.Write([]byte(rspStr))
		fmt.Printf("Server send msg:%s\n", rspStr)
		// time.Sleep(3 * time.Second)
	}
}
```
