package game

import (
	"fmt"
	"net"
	"os"
)

type Player struct {
	ID        string
	Name      string
	IPAddress string
	Port      int
	Score     int
	Connected bool
}

func (p *Player) Listen() {
	addr := net.UDPAddr{
		Port: p.Port,
		IP:   net.ParseIP(p.IPAddress),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	fmt.Printf("Listening on %s:%d\n", p.IPAddress, p.Port)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Printf("Received message from %v: %s\n", remoteAddr, string(buf[:n]))
	}
}

func (p *Player) SendMessage(to *Player) {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		Port: to.Port,
		IP:   net.ParseIP(to.IPAddress),
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := fmt.Sprintf("Hello from %s", p.Name)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Write error: ", err)
		os.Exit(1)
	}

	fmt.Println("Message sent:", msg)
}
