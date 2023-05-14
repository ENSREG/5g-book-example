package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	connNum := 100
	buf := make([]byte, 1024)

	wg.Add(connNum)
	for i := 0; i < connNum; i++ {
		go func() {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println("Error dialing:", err)
			}
			_, err = conn.Write([]byte("TCP Connection from client"))
			conn.Read(buf)
			conn.Close()
			wg.Done()
		}()

	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Nanosecond() - start.Nanosecond())
}
