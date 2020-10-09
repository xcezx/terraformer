package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NotificationGroupGenerator struct {
	MackerelService
}

func (g *NotificationGroupGenerator) InitResources() error {
	groups, err := g.Client().FindNotificationGroups()
	if err != nil {
		return err
	}
	for _, group := range groups {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			group.ID,
			fmt.Sprintf("notification_group_%s", group.ID),
			"mackerel_notification_group",
			"mackerel",
			[]string{},
		))
	}
	return nil
}
