package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mid/avg/avgpb"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := avgpb.NewCalcServiceClient(conn)
	doCalc(c)
}


func doCalc(c avgpb.CalcServiceClient) {

	requests := []*avgpb.CalculateRequest{
		{
			Calc: &avgpb.Calculate{
				Number: 4,
			},
		},
		{
			Calc: &avgpb.Calculate{
				Number: 5,
			},
		},
		{
			Calc: &avgpb.Calculate{
				Number: 4,
			},
		},
		{
			Calc: &avgpb.Calculate{
				Number: 6,
			},
		},
	}

	ctx := context.Background()
	stream, err := c.Calculate(ctx)
	if err != nil {
		log.Fatalf("error while calling CalculateAvg: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from Calculate: %v", err)
	}
	fmt.Printf("Average calculate Response: %v\n", res)
}
