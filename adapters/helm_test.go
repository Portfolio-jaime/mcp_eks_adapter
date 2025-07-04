package adapters

import (
	"testing"
)

func TestListHelmCharts(t *testing.T) {
	_, err := ListHelmCharts()
	if err == nil {
		t.Error("Se esperaba error al no estar en un entorno Helm real")
	}
}
