package cmd

import (
	"github.com/GoogleCloudPlatform/terraformer/providers/mackerel"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdMackerelImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mackerel",
		Short: "Import current state to Terraform configuration from Mackerel",
		Long:  "Import current state to Terraform configuration from Mackerel",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newMackerelProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newMackerelProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "service", "service=my-service")

	return cmd
}

func newMackerelProvider() terraformutils.ProviderGenerator {
	return &mackerel.MackerelProvider{}
}
