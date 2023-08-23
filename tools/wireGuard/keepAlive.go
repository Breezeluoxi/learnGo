package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverIP := "115.29.200.138"
	serverPort := 50814

	for {
		conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", serverIP, serverPort))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 发送虚拟数据
		_, err = conn.Write([]byte("heartbeat"))
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Heartbeat sent successfully")
		}

		conn.Close()

		// 等待一段时间，继续发送
		time.Sleep(25 * time.Second) // 与 PersistentKeepalive 保活间隔相同
	}
}
