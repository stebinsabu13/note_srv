package main

import (
	"fmt"
	"log"
	"net"

	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/config"
	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/db"
	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/pb"
	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed to load config", err)
	}
	h := db.Initdb(c)
	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)
	s := services.Server{
		H: h,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNoteServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
