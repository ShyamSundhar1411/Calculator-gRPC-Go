package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)

func displayPrimes(c pb.CalculatorServiceClient){
	var number int64;
	fmt.Println("Enter the Number")
	fmt.Scanln(&number)
	request:=&pb.PrimeRequest{
		Number: number,
	}
	stream,err:=c.PrimeGen(context.Background(),request)
	if err!=nil{
		log.Fatalf("Error Fetchimg Stuffs: %v\n",err)
	}else{
		for{
			message,err:=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("Error Fetching Stuffs: %v\n",err)
			}
			log.Println(message.Result)
		}
	}
}