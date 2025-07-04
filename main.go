package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Portfolio-jaime/mcp_eks_adapter/adapters"

	pb "github.com/Portfolio-jaime/mcp_eks_adapter/mcp/proto"
)

type server struct {
	pb.UnimplementedModelContextServiceServer
}

func (s *server) GetContext(ctx context.Context, req *pb.ContextRequest) (*pb.ContextResponse, error) {
	version, err := adapters.GetKubernetesVersion()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo versión k8s: %v", err)
	}
	charts, err := adapters.ListHelmCharts()
	if err != nil {
		return nil, fmt.Errorf("error listando releases helm: %v", err)
	}
	return &pb.ContextResponse{
		KubernetesVersion: version,
		Charts:            charts,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error escuchando: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterModelContextServiceServer(grpcServer, &server{})

	log.Println("Servidor MCP escuchando en :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("falló grpcServer: %v", err)
	}
}
