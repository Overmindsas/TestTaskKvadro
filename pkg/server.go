package pkg

import (
	"context"

	pb "github.com/Overmindsas/TestTaskKvadro/internal/proto"
)

type APIServer struct {
	pb.UnimplementedAPIServer
}

// функция для поиска автора по книге
func (a *APIServer) FindBook(ctx *context.Context, in *pb.Author) (*pb.Book, error) {

	return &pb.Book{Book: in.GetAuthor()}, nil
}

// функция для поиска книги по автору
func (a *APIServer) FindAuthor(ctx *context.Context, in *pb.Book) (*pb.Author, error) {

	return &pb.Author{Author: in.GetBook()}, nil
}
