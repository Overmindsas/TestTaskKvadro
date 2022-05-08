package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/Overmindsas/TestTaskKvadro/internal/proto"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type APIServer struct {
	pb.UnimplementedAPIServer
	db *sql.DB
}

func NewAPIServer() *APIServer {
	return &APIServer{db: &sql.DB{}}
}

// функция для поиска книги по автору
func (a *APIServer) FindBook(ctx context.Context, in *pb.Author) (*pb.BookList, error) {
	var book_list *pb.BookList = &pb.BookList{}

	rows, err := a.db.Query("SELECT * FROM BookAuthor WHERE author OR book = %s", string(in.Author))

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		b := pb.Book{}
		err := rows.Scan(&pb.Book{})
		if err != nil {
			return nil, err
		}
		book_list.Books = append(book_list.Books, &b)
	}

	return book_list, nil
}

// функция для поиска атора по книге
func (a *APIServer) FindAuthor(ctx context.Context, in *pb.Book) (*pb.Authorlist, error) {

	var author_list *pb.Authorlist = &pb.Authorlist{}

	rows, err := a.db.Query("SELECT * FROM BookAuthor WHERE author OR book = %s", string(in.Book))

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		au := pb.Author{}
		err := rows.Scan(&au.Author)
		if err != nil {
			return nil, err
		}
		author_list.Authors = append(author_list.Authors, &au)
	}

	return author_list, nil
}
func (server *APIServer) Run() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAPIServer(grpcServer, server)
	log.Printf("server listening at %v", lis.Addr())

	return grpcServer.Serve(lis)
}
func main() {
	var fff *APIServer = NewAPIServer()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8000)/kvadodb?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	fmt.Println(db.Ping())

	if err := fff.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
