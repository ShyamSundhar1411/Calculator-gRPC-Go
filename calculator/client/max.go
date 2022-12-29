package main

import(
	"context"
	"io"
	"log"
	"time"
	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)

func displayMax(c pb.CalculatorServiceClient){
	stream,err:=c.Max	(context.Background())
	if err!=nil{
		log.Fatalf("Error while connecting %v\n",err)
	}
	reqs:=[]*pb.MaxRequest{
		{Number:1},
		{Number:5},
		{Number:3},
		{Number:6},
		{Number:2},
		{Number:20},
	}
	waitChannel := make(chan struct{})
	go func(){
		for _,req := range reqs{
			log.Printf("Sending request %v\n",req)
			stream.Send(req)
			time.Sleep(1*time.Second)

		}
		stream.CloseSend()
	}()
	go func(){
		for{
			response,err:=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("Error while receiving %v\n",err)
				break
			}
			log.Println(response.Result)
		}
		close(waitChannel)
	}()
	<-waitChannel

}