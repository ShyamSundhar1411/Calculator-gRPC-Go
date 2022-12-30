package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server)SquareRoot(ctx context.Context, in *pb.SquareRootRequest) (*pb.SquareRootResponse, error){
	log.Printf("Invoked Square Root %v\n",in)
	responseNumber := in.Number
	if responseNumber<0{
		return nil,status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a Negative Number: %d",responseNumber),
		)
	}
	return &pb.SquareRootResponse{
		Result: math.Sqrt(float64(responseNumber)),
	},nil
}