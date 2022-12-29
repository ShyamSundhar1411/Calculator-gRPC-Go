package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)

func displayAverage(c pb.CalculatorServiceClient){
	requests := []*pb.AverageRequest{
		{Number:1},
		{Number:2},
		{Number:3},
		{Number:4},
	}
	stream,err:=c.Average(context.Background())
	if err!=nil{
		log.Fatalf("Error fetching: %v\n",err)
	}
	for _,request := range requests{
		log.Printf("Sending Request : %v\n",request)
		stream.Send(request)
		time.Sleep(1*time.Second)
	}
	response,err:=stream.CloseAndRecv()
	if err!=nil{
		log.Fatalf("Error fetching: %v\n",err)
	}else{
		log.Println(response)
	}
}

