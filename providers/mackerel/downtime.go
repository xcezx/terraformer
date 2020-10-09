package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DowntimeGenerator struct {
	MackerelService
}

func (g *DowntimeGenerator) InitResources() error {
	downtimes, err := g.Client().FindDowntimes()
	if err != nil {
		return err
	}
	for _, downtime := range downtimes {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			downtime.ID,
			fmt.Sprintf("downtime_%s", downtime.ID),
			"mackerel_downtime",
			"mackerel",
			[]string{},
		))
	}
	return nil
}
