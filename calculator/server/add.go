package main

import(
	"context"
	"log"
	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)
func (s *Server)Add(ctx context.Context,in *pb.AddRequest)(*pb.AddResponse,error){
	log.Printf("Add Function was invoked with %v\n",in)
	result:=in.FirstNumber+in.SecondNumber
	return &pb.AddResponse{
		Result:result,
	},nil
}