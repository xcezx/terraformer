package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

type ServiceGenerator struct {
	MackerelService
}

func (g *ServiceGenerator) InitResources() error {
	client := g.Client()

	services, err := client.FindServices()
	if err != nil {
		return err
	}
	for _, service := range services {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			service.Name,
			fmt.Sprintf("service_%s", service.Name),
			"mackerel_service",
			g.ProviderName,
			[]string{},
		))

		roleResources, err := g.createRoleResources(client, service)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, roleResources...)
	}
	return nil
}

func (g *ServiceGenerator) createRoleResources(client *mackerel.Client, service *mackerel.Service) ([]terraformutils.Resource, error) {
	roles, err := client.FindRoles(service.Name)
	if err != nil {
		return nil, err
	}
	var resources []terraformutils.Resource
	for _, role := range roles {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s:%s", service.Name, role.Name),
			fmt.Sprintf("role_%s_%s", service.Name, role.Name),
			"mackerel_role",
			g.ProviderName,
			map[string]string{
				"service": service.Name,
				"name":    role.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources, nil
}

func (g *ServiceGenerator) PostConvertHook() error {
	for _, service := range g.Resources {
		if service.InstanceInfo.Type != "mackerel_service" {
			continue
		}
		for i, role := range g.Resources {
			if role.InstanceInfo.Type != "mackerel_role" {
				continue
			}
			if role.InstanceState.Attributes["service"] == service.InstanceState.Attributes["name"] {
				g.Resources[i].Item["service"] = "${mackerel_service." + service.ResourceName + ".name}"
			}
		}
	}
	return nil
}
