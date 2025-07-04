package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"

	pb "./proto"
)

type server struct {
	pb.UnimplementedModelContextServiceServer
}

func (s *server) GetContext(ctx context.Context, req *pb.ContextRequest) (*pb.ContextResponse, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("error creando config in-cluster: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creando cliente kubernetes: %v", err)
	}

	versionInfo, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo versión k8s: %v", err)
	}

	helmSettings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(helmSettings.RESTClientGetter(), "", "secrets", log.Printf); err != nil {
		return nil, fmt.Errorf("error inicializando helm config: %v", err)
	}

	list := action.NewList(actionConfig)
	list.All = true
	releases, err := list.Run()
	if err != nil {
		return nil, fmt.Errorf("error listando releases helm: %v", err)
	}

	var charts []*pb.HelmChart
	for _, rel := range releases {
		charts = append(charts, &pb.HelmChart{
			Name:      rel.Name,
			Version:   rel.Chart.Metadata.Version,
			Namespace: rel.Namespace,
		})
	}

	return &pb.ContextResponse{
		KubernetesVersion: versionInfo.GitVersion,
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
