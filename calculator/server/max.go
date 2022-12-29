package main

import (
	"io"
	"log"
	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)

func (s *Server)Max(stream pb.CalculatorService_MaxServer) error{
	log.Panicln("Invoked Max Function")
	maximum := int64(0)
	for{
		request,err:=stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err!=nil{
			log.Fatalf("Error while reading client stream %v\n",err)
		}
		response := request.Number
		if(response>maximum){
			maximum = response
		}
		err = stream.Send(&pb.MaxResponse{
			Result: maximum,
		})
		if err!=nil{
			log.Fatalf("Error sending data to client %v\n",err)
		}
	}	
}