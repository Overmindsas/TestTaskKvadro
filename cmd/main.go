package main

import (
	"log"
	"net"

	pb "github.com/Overmindsas/TestTaskKvadro/internal/proto"
	"github.com/Overmindsas/TestTaskKvadro/pkg"
	"google.golang.org/grpc"
)

func main() {
	// инициализация grpc сервера
	grpcServer := grpc.NewServer()

	// инициализация нашей структурки)
	srv := &pkg.APIServer{}

	// второй аргумент чёт я нихуя не понял
	pb.RegisterAPIServer(grpcServer, srv.UnimplementedAPIServer)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
