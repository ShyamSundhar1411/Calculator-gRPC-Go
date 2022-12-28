package main
import(
	"context"
	"log"
	"fmt"
	pb "github.com/ShyamSundhar1411/Calculator-gRPC/calculator/proto"
)
func addFunction(c pb.AddServiceClient){
	log.Println("Add Function was invoked")
	var firstNumber int64
	var secondNumber int64
	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)
	response,err:=c.Add(context.Background(),&pb.AddRequest{
		FirstNumber: firstNumber,
		SecondNumber: secondNumber,
	})
	if err!=nil{
		log.Fatalf("Could not add: %v\n",err)
	}
	log.Printf("Result : %d\n",response.Result)
}