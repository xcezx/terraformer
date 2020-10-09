package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ChannelGenerator struct {
	MackerelService
}

func (g *ChannelGenerator) InitResources() error {
	channels, err := g.Client().FindChannels()
	if err != nil {
		return err
	}
	for _, channel := range channels {
		switch channel.Type {
		case "email", "slack", "webhook":
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				channel.ID,
				fmt.Sprintf("channel_%s", channel.ID),
				"mackerel_channel",
				"mackerel",
				[]string{},
			))
		}
	}
	return nil
}
