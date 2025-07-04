package adapters

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// GetKubernetesVersion obtiene la versión del cluster Kubernetes
func GetKubernetesVersion() (string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}
	versionInfo, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}
	return versionInfo.GitVersion, nil
}
