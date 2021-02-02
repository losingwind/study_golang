package main

import (
	"log"
	"net"
	"net/rpc"
	"study_golang/studyGrpc/go-rpc/pkg"
)

type Server struct{}

func (s *Server) Server(request pkg.String, rsp *pkg.String) error {
	rsp.Value = request.Value
	return nil
}

func main() {
	if err := pkg.Register(new(Server)); err != nil {
		log.Fatal(err.Error())
	}

	listener, err := net.Listen("tcp", pkg.Addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	rpc.Accept(listener)
}
