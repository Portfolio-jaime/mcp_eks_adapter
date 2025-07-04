# MCP EKS Adapter

Este proyecto implementa un servidor gRPC compatible con el Model Context Protocol (MCP) que se conecta a un cluster de Kubernetes (EKS) y devuelve información relevante como:

- Versión de Kubernetes
- Charts instalados con Helm

## Estructura

- `main.go`: Servidor gRPC principal
- `adapters/`: Lógica para interactuar con Kubernetes y Helm
- `mcp/proto/`: Definición y código generado de Protobuf
- `client/client.go`: Cliente de ejemplo para probar el servicio

## Requisitos

- Go 1.21+
- Helm instalado
- Acceso al cluster EKS (desde dentro del clúster o con kubeconfig adecuado)
- Protoc y plugins para Go (`protoc-gen-go`, `protoc-gen-go-grpc`)

## Compilación y ejecución

### 1. Compilar el servidor

```sh
go build -o mcp-eks-adapter main.go
```

### 2. Ejecutar el servidor

```sh
./mcp-eks-adapter
```

### 3. Probar con el cliente

En otra terminal:

```sh
go run client/client.go
```

## Probar con un cluster EKS real

- El servidor debe ejecutarse en un entorno con acceso al API de Kubernetes (por ejemplo, dentro de un pod en EKS o con kubeconfig local configurado).
- El ServiceAccount (si es dentro de EKS) debe tener permisos de lectura de recursos y Helm.
- El cliente mostrará la versión real de Kubernetes y los charts instalados con Helm en el cluster.

## Notas

- Si ejecutas localmente, asegúrate de tener acceso a tu cluster EKS con `kubectl config use-context ...`.
- Puedes modificar el cliente para enviar diferentes nombres de cluster si lo deseas.

---

¿Dudas? Abre un issue o contacta al autor.
