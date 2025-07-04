package adapters

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetKubernetesVersion obtiene la versión del cluster Kubernetes
func GetKubernetesVersion() (string, error) {
	var config *rest.Config
	var err error

	fmt.Println("[DEBUG] Entrando a GetKubernetesVersion")
	fmt.Printf("[DEBUG] KUBERNETES_SERVICE_HOST=%s\n", os.Getenv("KUBERNETES_SERVICE_HOST"))
	fmt.Printf("[DEBUG] KUBERNETES_SERVICE_PORT=%s\n", os.Getenv("KUBERNETES_SERVICE_PORT"))

	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		fmt.Println("[DEBUG] Usando InClusterConfig")
		config, err = rest.InClusterConfig()
	} else {
		fmt.Println("[DEBUG] Usando kubeconfig local")
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		fmt.Printf("[DEBUG] Error obteniendo config: %v\n", err)
		return "", err
	}

	fmt.Println("[DEBUG] Creando clientset de Kubernetes")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("[DEBUG] Error creando clientset: %v\n", err)
		return "", err
	}

	fmt.Println("[DEBUG] Obteniendo versión del servidor Kubernetes")
	versionInfo, err := clientset.Discovery().ServerVersion()
	if err != nil {
		fmt.Printf("[DEBUG] Error obteniendo versión: %v\n", err)
		return "", err
	}
	fmt.Printf("[DEBUG] Versión obtenida: %s\n", versionInfo.GitVersion)
	return versionInfo.GitVersion, nil
}
