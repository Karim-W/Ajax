package cmd

import (
	"github.com/karim-w/Ajax/services/codegensvc"
	"github.com/spf13/cobra"
)

var GenCmd = &cobra.Command{
	Use:     "g",
	Aliases: []string{"gen"},
	Short:   "Generate Code",
	Long:    `Generate Code`,
	Run: func(cmd *cobra.Command, args []string) {
		commands := make(map[string]string, 3)
		controller := cmd.Flag("controller").Value.String()
		service := cmd.Flag("service").Value.String()
		router := cmd.Flag("router").Value.String()
		index := cmd.Flag("index").Value.String()
		commands["controller"] = controller
		commands["service"] = service
		commands["router"] = router
		commands["index"] = index
		codegensvc.CodeGenerator(commands)
	},
}

func init() {
	RootCmd.AddCommand(GenCmd)
	GenCmd.Flags().StringP("controller", "c", "", "Generate Api Controller")
	GenCmd.Flags().StringP("service", "s", "", "Generate Api Service")
	GenCmd.Flags().StringP("router", "r", "", "Generate Api Router")
	GenCmd.Flags().StringP("index", "i", "", "Generate Index Page")
}
