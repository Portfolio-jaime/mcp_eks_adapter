package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Portfolio-jaime/mcp_eks_adapter/mcp/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewModelContextServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	resp, err := client.GetContext(ctx, &pb.ContextRequest{ClusterName: "test"})
	if err != nil {
		log.Fatalf("Error llamando GetContext: %v", err)
	}

	log.Printf("Versión de Kubernetes: %s", resp.KubernetesVersion)
	for _, chart := range resp.Charts {
		log.Printf("Chart: %s, Version: %s, Namespace: %s", chart.Name, chart.Version, chart.Namespace)
	}
}
