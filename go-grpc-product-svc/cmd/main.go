package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Akshayvij07/go-grpc-project/gp-grpc-product-svc/pkg/db"
	services "github.com/Akshayvij07/go-grpc-project/gp-grpc-product-svc/pkg/services"
)

func main() {
	c, err := conifg.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
