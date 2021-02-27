package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
	"mid/calc/calcpb"
	"net"
)

type Server struct {
	calcpb.UnimplementedCalcServiceServer
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

func (s *Server) Calculate(req *calcpb.CalculateRequest, stream calcpb.CalcService_CalculateServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v \n", req)
	number := req.GetCalc().GetNumber()

	for number%2 == 0 {
		res := &calcpb.CalculateResponse{Res: 2}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending result: %v", err.Error())
		}
		number/=2
	}

	for i := 3; float64(i) <= math.Sqrt(float64(number)); i+=2 {
		for int(number) % i ==0 {
			res := &calcpb.CalculateResponse{Res: int64(i)}
			if err := stream.Send(res); err != nil {
				log.Fatalf("error while sending result: %v", err.Error())
			}
			number/=int64(i)
		}
	}

	if number > 2{
		res := &calcpb.CalculateResponse{Res: int64(number)}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending result: %v", err.Error())
		}
	}

	return nil
}
