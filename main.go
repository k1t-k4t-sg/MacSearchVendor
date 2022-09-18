package main

import (
	"log"
	"net"

	pb "MacSearchVendor/pkg/api"
	mac "MacSearchVendor/pkg/mac"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterSearchVendorServer(s, &mac.GrpcServer{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
