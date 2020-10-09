package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

type MetadataGenerator struct {
	MackerelService
}

func (g *MetadataGenerator) InitResources() error {
	client := g.Client()
	services, err := client.FindServices()
	if err != nil {
		return err
	}
	for _, service := range services {
		serviceMetadataResources, err := g.createServiceMetadataResources(client, service)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, serviceMetadataResources...)

		roles, err := client.FindRoles(service.Name)
		if err != nil {
			return err
		}
		for _, role := range roles {
			roleMetadataResources, err := g.createRoleMetadataResources(client, service, role)
			if err != nil {
				return err
			}
			g.Resources = append(g.Resources, roleMetadataResources...)
		}
	}
	return nil
}

func (g *MetadataGenerator) createServiceMetadataResources(client *mackerel.Client, service *mackerel.Service) ([]terraformutils.Resource, error) {
	namespaces, err := client.GetServiceMetaDataNameSpaces(service.Name)
	if err != nil {
		return nil, err
	}
	var resources []terraformutils.Resource
	for _, namespace := range namespaces {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s/%s", service.Name, namespace),
			fmt.Sprintf("service_metadata_%s/%s", service.Name, namespace),
			"mackerel_service_metadata",
			g.ProviderName,
			map[string]string{
				"service":   service.Name,
				"namespace": namespace,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources, nil
}

func (g *MetadataGenerator) createRoleMetadataResources(client *mackerel.Client, service *mackerel.Service, role *mackerel.Role) ([]terraformutils.Resource, error) {
	namespaces, err := client.GetRoleMetaDataNameSpaces(service.Name, role.Name)
	if err != nil {
		return nil, err
	}
	var resources []terraformutils.Resource
	for _, namespace := range namespaces {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s:%s/%s", service.Name, role.Name, namespace),
			fmt.Sprintf("role_metadata_%s:%s/%s", service.Name, role.Name, namespace),
			"mackerel_role_metadata",
			g.ProviderName,
			map[string]string{
				"service":   service.Name,
				"role":      role.Name,
				"namespace": namespace,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources, nil
}
