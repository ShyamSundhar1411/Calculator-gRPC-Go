package main

import (
	"io"
	"log"

	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error{
	var addedSum int
	counter:=0
	for{
		request,err := stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&pb.AverageResponse{
				Result : float64(addedSum/counter),
			})
		}
		if err!=nil{
			log.Fatalf("Error Fetching : %v\n",err)
		}else{
			addedSum+=int(request.Number)
			counter++
		}
	}
}