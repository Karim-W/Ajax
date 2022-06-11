package cmd

import (
	"github.com/karim-w/Ajax/services/clientsvc"
	"github.com/spf13/cobra"
)

var ClientCmd = &cobra.Command{
	Use:     "c",
	Aliases: []string{"client"},
	Short:   "Generate Client",
	Long:    `Generate Client based on openapi.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		commands := make(map[string]string, 3)
		language := cmd.Flag("language").Value.String()
		path := cmd.Flag("path").Value.String()
		commands["language"] = language
		commands["path"] = path
		clientsvc.ClientGenerator(commands)

	},
}

func init() {
	RootCmd.AddCommand(GenCmd)
	GenCmd.Flags().StringP("language", "l", "go", "Generate Code")
	GenCmd.Flags().StringP("path", "p", "", "Generate Code")
}
