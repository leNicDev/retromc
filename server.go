package main

import (
	"log"
	"net"

	"github.com/leNicDev/retromc/packethandler"
)

const (
	CON_HOST = "localhost"
	CON_PORT = "25565"
	CON_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CON_TYPE, CON_HOST+":"+CON_PORT)
	if err != nil {
		log.Panicln("Failed to bind to address", err.Error())
	}

	// close listener when the application closes
	defer l.Close()

	log.Println("Server listening on " + CON_HOST + ":" + CON_PORT)

	for {
		// listen for incoming connections
		connection, err := l.Accept()
		if err != nil {
			log.Fatalln("Failed to accept connection: ", err.Error())
			continue
		}

		// handle connection
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	// create buffer to hold incoming data
	var buf []byte

	for {
		buf = make([]byte, 1024)

		// read incoming data
		_, err := connection.Read(buf)
		if err != nil {
			log.Fatalln("Failed to read packet: ", err.Error())
			return
		}

		packethandler.HandlePacket(connection, &buf)
	}
}
