package cmd

import (
	"github.com/karim-w/Ajax/services/dockersvc"
	"github.com/spf13/cobra"
)

var DockerCmd = &cobra.Command{
	Use:     "d",
	Aliases: []string{"docker"},
	Short:   "Manage Docker Containers",
	Long:    `Manage Docker Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		commands := make(map[string]string, 3)
		list := cmd.Flag("list").Value.String()
		pull := cmd.Flag("get").Value.String()
		kill := cmd.Flag("kill").Value.String()
		pregen := cmd.Flag("pregen").Value.String()
		commands["list"] = list
		commands["get"] = pull
		commands["kill"] = kill
		commands["pregen"] = pregen
		dockersvc.DockerService(commands)
	},
}

func init() {
	RootCmd.AddCommand(DockerCmd)
	DockerCmd.Flags().BoolP("list", "l", false, "List Docker Containers")
	DockerCmd.Flags().StringP("get", "g", "", "Pull Docker Container")
	DockerCmd.Flags().StringP("kill", "k", "", "Kill Docker Container")
	DockerCmd.Flags().StringP("pregen", "p", "", "Pre-generated Containers")
}
