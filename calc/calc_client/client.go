package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"mid/calc/calcpb"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)
	doCalc(c)
}


func doCalc(c calcpb.CalcServiceClient) {

	ctx := context.Background()
	req := &calcpb.CalculateRequest{Calc : &calcpb.Calculate{
		Number: 200,
	}}

	stream, err := c.Calculate(ctx, req)
	if err != nil {
		log.Fatalf("error while calling Calculate RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from PrimeNumberDecomposition RPC %v", err)
		}
		log.Printf("response from PrimeNumberDecomposition:%v \n", res.GetRes())
	}
}
