package mackerel

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
	"github.com/zclconf/go-cty/cty"
)

type MackerelProvider struct { //nolint
	terraformutils.Provider
	apiKey string
}

func (p *MackerelProvider) Init(_ []string) error {
	apiKey := os.Getenv("MACKEREL_API_KEY")
	if apiKey == "" {
		return errors.New("set MACKEREL_API_KEY env var")
	}
	p.apiKey = apiKey

	return nil
}

func (p *MackerelProvider) GetName() string {
	return "mackerel"
}

func (p *MackerelProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal(p.apiKey),
	})
}

func (p *MackerelProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_key": p.apiKey,
	})

	return nil
}

func (p *MackerelProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"alert_group_setting": &AlertGroupSettingGenerator{},
		"channel":             &ChannelGenerator{},
		"downtime":            &DowntimeGenerator{},
		"metadata":            &MetadataGenerator{},
		"monitor":             &MonitorGenerator{},
		"notification_group":  &NotificationGroupGenerator{},
		"service":             &ServiceGenerator{},
	}
}

func (p *MackerelProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *MackerelProvider) GetProviderData(_ ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"version": providerwrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}
