package main

import (
	"fmt"
	"syscall"
)

type Socket struct {
	skt    int
	closed bool
}

func (s *Socket) CreateServer(port int) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Socket created with fd:", fd)

	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: port})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = syscall.Listen(fd, 20)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.skt = fd
	s.closed = false
}

func (s *Socket) Accept() (Socket, error) {
	fd, _, err := syscall.Accept(s.skt)
	if err != nil {
		fmt.Println(err)
		return Socket{}, err
	}
	return Socket{fd, false}, nil
}

func (s *Socket) Close() {
	if s.closed {
		return
	}
	syscall.Close(s.skt)
	s.closed = true
}

func (s *Socket) recv_some(was_closed *bool) {
	buff := make([]byte, 1024)
	n, err := syscall.Read(s.skt, buff)
	if err != nil {
		*was_closed = true
		fmt.Println(err)
		return
	}

	if n == 0 {
		*was_closed = true
		return
	}

	fmt.Println("Received: ", string(buff[:n]))
}

func main() {
	s := &Socket{}
	s.CreateServer(8080)
	new_connection, err := s.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	was_closed := false
	for was_closed == false {
		new_connection.recv_some(&was_closed)
	}

	s.Close()
}
