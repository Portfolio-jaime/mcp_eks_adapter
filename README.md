# MCP EKS Adapter

Este proyecto implementa un servidor gRPC compatible con el Model Context Protocol (MCP) que se conecta a un cluster de Kubernetes (EKS) y devuelve información relevante como:

- Versión de Kubernetes
- Charts instalados con Helm

## Requisitos

- Go 1.21+
- Helm instalado
- Acceso al cluster (desde dentro del clúster con ServiceAccount adecuado)

## Compilación de Protobuf

```bash
cd proto
protoc --go_out=. --go-grpc_out=. context.proto
