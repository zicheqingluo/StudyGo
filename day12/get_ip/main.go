package main

import (
	"fmt"
	"net"
	"strings"
)

// GetOutboundIP 获取本地对外IP
func main()  {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip := strings.Split(localAddr.IP.String(), ":")[0]
	fmt.Println(ip)
	return
}