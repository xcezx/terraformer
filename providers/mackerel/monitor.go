package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type MonitorGenerator struct {
	MackerelService
}

func (g *MonitorGenerator) InitResources() error {
	monitors, err := g.Client().FindMonitors()
	if err != nil {
		return err
	}
	for _, monitor := range monitors {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			monitor.MonitorID(),
			fmt.Sprintf("monitor_%s", monitor.MonitorID()),
			"mackerel_monitor",
			"mackerel",
			[]string{},
		))
	}
	return nil
}
