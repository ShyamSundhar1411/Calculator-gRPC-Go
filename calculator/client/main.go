package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)

var address string = "localhost:5501"

func main() {
	connection,err:=grpc.Dial(address,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Failed to connect: %v\n",err)
	}
	defer connection.Close()
	client := pb.NewCalculatorServiceClient(connection)
	addFunction(client)
	displayPrimes(client)
	displayAverage(client)
	displayMax(client)
	displaySqrt(client,25)
	displaySqrt(client,-1)
}