package adapters

import (
	"testing"
)

func TestGetKubernetesVersion(t *testing.T) {
	_, err := GetKubernetesVersion()
	if err == nil {
		t.Error("Se esperaba error al no estar en un cluster Kubernetes real")
	}
}
