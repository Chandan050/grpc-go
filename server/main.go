package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/pascalallen/grpc-go/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "HEllo , world"}, nil
}

type Logindetails struct {
	Password string
	LoginAt  time.Time
	Logout   time.Time
}

var usercreated = make(map[string]*Logindetails)

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.HelloWorldResponse, error) {
	usercreated[in.Name] = &Logindetails{Password: in.Password}
	return &pb.HelloWorldResponse{Message: "User is created"}, nil
}

func (s *server) LoginUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.HelloWorldResponse, error) {
	user, exits := usercreated[in.GetName()]
	if !exits {
		return &pb.HelloWorldResponse{Message: "user doesn't exits"}, nil
	}
	if user.Password != in.GetPassword() {
		return &pb.HelloWorldResponse{Message: "incorrect password"}, nil
	}
	user.LoginAt = time.Now()
	return &pb.HelloWorldResponse{Message: "user login credential ar correct"}, nil
}

func (s *server) LogoutUser(ctx context.Context, in *pb.LogoutRequest) (*pb.HelloWorldResponse, error) {
	user, exits := usercreated[in.GetName()]
	if !exits {
		return &pb.HelloWorldResponse{Message: "user doesn't exits"}, nil
	}
	user.Logout = time.Now()
	return &pb.HelloWorldResponse{Message: "user logouted"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloWorldServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
