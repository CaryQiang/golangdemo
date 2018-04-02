package worker

import (
	"context"
	"fmt"
	pb "golangdemo/masterWorkerDemo/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

//TargetServer 目标服务
type TargetServer struct {
	config     *Config
	grpcServer *grpc.Server
	stop       chan bool
	shutdown   chan bool
}

//GetTargetServer create TargetServer
func GetTargetServer(conf *Config) (*TargetServer, error) {
	var serverOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(serverOptions...)

	s := TargetServer{
		config:     conf,
		grpcServer: grpcServer,
		stop:       make(chan bool, 1),
	}

	pb.RegisterTargetServiceServer(grpcServer, &s)

	return &s, nil
}

//Start start http server
func (s *TargetServer) Start() error {
	lis, err := net.Listen("tcp", s.config.TargetGrpcpoint)
	if err != nil {
		log.Println("Failed to listen tcp server: ", err)
		panic(err)
	}

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			log.Println("Failed to start grpc server: ", err)
			panic(err)
		}
	}()
	log.Printf("Target Server: grpc server listing at %s...\n", s.config.TargetGrpcpoint)

	go s.Run()

	return nil
}

func (s *TargetServer) Run() {

	for {
		select {
		case <-s.stop:
			goto OnCloseLabel
		}
	}
OnCloseLabel:
	close(s.shutdown)
	log.Println("Target Server: active server is shutting down")
}

func (s *TargetServer) Stop() {
	s.stop <- true

	<-s.shutdown

	log.Println("Target Server is shut-down")
}

func (s *TargetServer) TargetLibList(ctx context.Context, in *pb.TargetLibListRequest) (*pb.TargetLibListResponse, error) {
	//uuid
	res := pb.TargetLibListResponse{Libs: make([]*pb.TargetLib, 0)}

	res.Libs = append(res.Libs, &pb.TargetLib{
		LibId:       "00000000",
		Name:        "test Lib",
		Type:        0,
		TargetCount: 1,
	})

	fmt.Println("TargetLibList return...")

	return &res, nil
}

func (s *TargetServer) TargetAdd(ctx context.Context, in *pb.TargetAddRequest) (*pb.TargetAddResponse, error) {
	fmt.Printf("TargetAdd received data: %v\n", in.Target)

	return &pb.TargetAddResponse{TargetId: "0000"}, nil
}

func (s *TargetServer) TargetDelete(ctx context.Context, in *pb.TargetDelRequest) (*pb.TargetDelResponse, error) {

	fmt.Printf("TargetDelete received data: %s\n", in.TargetId)
	return &pb.TargetDelResponse{}, nil
}
