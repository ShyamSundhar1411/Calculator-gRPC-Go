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
			result:=float64(addedSum)/float64(counter)
			return stream.SendAndClose(&pb.AverageResponse{
				Result : result,
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