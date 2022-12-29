package main

import (
	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)
func (s *Server)PrimeGen(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeGenServer) error{
	k:=int64(2)
	number := in.Number
	for number>1{
		if number%k==0{
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			number/=k
		}else{
			k=k+1
		}
	}
	return nil
}