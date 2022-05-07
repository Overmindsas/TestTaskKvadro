package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Overmindsas/TestTaskKvadro/internal/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	client := pb.NewAPIClient(conn)

	msg, err := client.FindAuthor(context.Background(), &pb.Book{Book: "1984"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(msg.GetAuthor())
}
