package mackerel

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

type MackerelService struct { //nolint
	terraformutils.Service
}

func (s *MackerelService) Client() *mackerel.Client {
	client := mackerel.NewClient(s.Args["api_key"].(string))
	client.Verbose = s.Verbose
	return client
}
