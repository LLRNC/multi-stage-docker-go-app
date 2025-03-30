package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    version := os.Getenv("VERSION")
    if version == "" {
        version = "unknown"
    }

    ip := getOutboundIP()
    hostname, _ := os.Hostname()

    fmt.Printf("IP serwera: %s\n", ip)
    fmt.Printf("Hostname: %s\n", hostname)
    fmt.Printf("Wersja aplikacji: %s\n", version)
}

func getOutboundIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        return "unknown"
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP.String()
}
