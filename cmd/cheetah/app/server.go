package app

import (
	"github.com/spf13/cobra"

	"cheetah/cmd/cheetah/app/options"
)

func NewServerCommand() *cobra.Command {
	o := options.NewAppOptions()
	cmd := &cobra.Command{
		Use:  "cheetah",
		Long: `Long describe.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(o)
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of cheetah.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("0.0.7")
		},
	}
	projectCmd := &cobra.Command{
		Use:   "create",
		Short: "Create [type] project.",
		Long:  "Create [type] project. type is 'ansible'、'gitbook'、'mdbook'、'command'、'http'、'mvc'、'grpc'、'simple'.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Println("Please specify the project item.")
			}
		},
	}
	projectCmd.Flags().StringVarP(&o.Name, "name", "n", "", "Project name.")
	projectCmd.Flags().StringVarP(&o.Path, "path", "p", "", "Project path.")

	cmd.AddCommand(versionCmd)
	cmd.AddCommand(projectCmd)

	return cmd
}

func run(o *options.AppOptions) (err error) {
	server, err := o.NewServer()
	server.Run()
	return
}
