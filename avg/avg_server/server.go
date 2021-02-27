package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"mid/avg/avgpb"
	"net"
)

type Server struct {
	avgpb.UnimplementedCalcServiceServer
}


func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	avgpb.RegisterCalcServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

func (s *Server) Calculate(stream avgpb.CalcService_CalculateServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	nums := []int64{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			var result float64
			sum := 0
			for i:=0; i< len(nums); i++{
				sum += int(nums[i])
			}
			result = float64(sum) / float64(len(nums))
			return stream.SendAndClose(&avgpb.CalculateResponse{
				Res: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		nums = append(nums, req.Calc.GetNumber())
	}
}

