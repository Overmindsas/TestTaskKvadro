package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/Overmindsas/TestTaskKvadro/internal/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

func main() {

	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 1 {
		log.Fatal("not arg")
	}

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	client := pb.NewAPIClient(conn)

	msg, err := client.FindAuthor(context.Background(), &pb.Book{Book: flag.Arg(0)})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(msg.Authors)
}
