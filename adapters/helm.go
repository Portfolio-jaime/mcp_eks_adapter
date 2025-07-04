package adapters

import (
	"log"

	pb "github.com/Portfolio-jaime/mcp_eks_adapter/mcp/proto"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

// ListHelmCharts lista los charts instalados con Helm
func ListHelmCharts() ([]*pb.HelmChart, error) {
	helmSettings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(helmSettings.RESTClientGetter(), "", "secrets", log.Printf); err != nil {
		return nil, err
	}
	list := action.NewList(actionConfig)
	list.All = true
	releases, err := list.Run()
	if err != nil {
		return nil, err
	}
	var charts []*pb.HelmChart
	for _, rel := range releases {
		charts = append(charts, &pb.HelmChart{
			Name:      rel.Name,
			Version:   rel.Chart.Metadata.Version,
			Namespace: rel.Namespace,
		})
	}
	return charts, nil
}
