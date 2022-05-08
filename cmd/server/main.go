package main

import (
	"context"
	"log"
	"net"

	"github.com/Overmindsas/TestTaskKvadro/internal/db"
	pb "github.com/Overmindsas/TestTaskKvadro/internal/proto"
	"google.golang.org/grpc"
)

type APIServer struct {
	pb.UnimplementedAPIServer
}

// функция для поиска книги по автору
func (a *APIServer) FindBook(ctx context.Context, in *pb.Author) (*pb.Book, error) {

	return &pb.Book{}, nil
}

// функция для поиска атора по книге
func (a *APIServer) FindAuthor(ctx context.Context, in *pb.Book) (*pb.Author, error) {

	return &pb.Author{}, nil
}

func main() {

	db, err := db.DbConnect()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	// инициализация grpc сервера
	grpcServer := grpc.NewServer()

	// инициализация нашей структурки
	//srv := &APIServer{}

	pb.RegisterAPIServer(grpcServer, &APIServer{})

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
