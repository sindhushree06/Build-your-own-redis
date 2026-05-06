package server

import (
	"io"
	"log"
	"net"
	"servers/config"
	"strconv"
)

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	e, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:e]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func RunSyncTCPServer() {
	log.Println("starting a synchronous TCP server on", config.Host, config.Port)
	var con_clients int = 0
	lsmr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
	for {
		c, err := lsmr.Accept()
		if err != nil {
			panic(err)
		}
		con_clients += 1
		log.Println("client connected with address:", c.RemoteAddr(), "concurrrent client", con_clients)
		for {
			cmd, err := readCommand(c)
			if err != nil {
				con_clients -= 1
				log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
				if err == io.EOF {
					break
				} else {
					c.Close()
				}
				log.Println("err", err)
			}
			log.Println(cmd)
			if err = respond(cmd, c); err != nil {
				log.Print("err write:", err)
			}
		}
	}
}
