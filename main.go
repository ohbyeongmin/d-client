package main

import (
	"context"
	"log"

	pb "github.com/ohbyeongmin/daangn-grpc-client/g-vote"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("not connect")
	}
	defer conn.Close()
	c := pb.NewVoteClient(conn)
	md := metadata.Pairs(
		"x-user-id", "qqqq",
	)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	// boardId := uint64(1111111)
	// title := "grpc"
	// description := "Grpc"
	// voteItems := []string{"g", "r", "p", "c"}
	// r, err := c.CreateVote(mdCtx, &pb.CreateQuery{BoardId: boardId, Title: title, Description: description, Vote_Items: voteItems})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(r)

	// r, err := c.SearchVote(mdCtx, &pb.SearchQuery{BoardId: 123456, VoteId: "6082226703"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(r)

	r, err := c.SelectVote(mdCtx, &pb.SelectQuery{BoardId: 3242212, VoteId: "6082226899", VoteItemId: 2})
	log.Println(r)
	if err != nil {
		log.Fatal(err)
	}
}
