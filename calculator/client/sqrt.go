package main

import (
	"context"
	"log"

	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func displaySqrt(c pb.CalculatorServiceClient,n int64){
	log.Println("Sqrt was invoked from Client Side")
	response,err:=c.SquareRoot(context.Background(),&pb.SquareRootRequest{Number: n})
	if err!=nil{
		e,ok:=status.FromError(err)
		if ok{
			log.Println("Error message from server: "+e.Message())
			log.Println("Error Code: "+e.Code().String())
			if e.Code() == codes.InvalidArgument{
				log.Println("Probably sent a negative number")
				return
			}

		}else{
			log.Fatalf("A non gRPC Error:%v\n",err)
		}
	}
	log.Printf("Sqrt: %f\n",response.Result)
}