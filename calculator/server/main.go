package main

import (
	"log"
	"net"

	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
	"google.golang.org/grpc"
)
var address string = "0.0.0.0:5501"

type Server struct{
	pb.CalculatorServiceServer
}

func main(){
	listener,err := net.Listen("tcp",address)
	if err!=nil{
		log.Fatalf("Failed listening on %v\n",err)
	}
	log.Printf("Listening on %s\n",address)
	serverInstance := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(serverInstance,&Server{})
	if err=serverInstance.Serve(listener);err!=nil{
		log.Fatalf("Failed to listen on %v\n",err)
	}
}