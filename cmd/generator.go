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
		commands["controller"] = controller
		codegensvc.CodeGenerator(commands)
	},
}

func init() {
	RootCmd.AddCommand(GenCmd)
	GenCmd.Flags().StringP("controller", "c", "", "Generate Api Controller")
}
