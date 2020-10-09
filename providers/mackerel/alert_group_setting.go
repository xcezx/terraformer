package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AlertGroupSettingGenerator struct {
	MackerelService
}

func (g *AlertGroupSettingGenerator) InitResources() error {
	settings, err := g.Client().FindAlertGroupSettings()
	if err != nil {
		return err
	}
	for _, setting := range settings {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			setting.ID,
			fmt.Sprintf("alert_group_setting_%s", setting.ID),
			"mackerel_alert_group_setting",
			"mackerel",
			[]string{},
		))
	}
	return nil
}
