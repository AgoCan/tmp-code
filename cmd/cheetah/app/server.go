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
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		DisableAutoGenTag: true,
	}
	cmd.CompletionOptions.HiddenDefaultCmd = true

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of cheetah.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("0.0.7")
		},
	}
	projectCmd := &cobra.Command{
		Use:   "create",
		Short: "Create type project.",
		Long:  "Create type project. type is 'ansible'、'gitbook'、'mdbook'、'command'、'http'、'mvc'、'grpc'、'simple'.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(projectCmd)

	projectCmd.PersistentFlags().StringVarP(&o.Name, "name", "n", "demo", "Project name.")
	projectCmd.PersistentFlags().StringVarP(&o.Path, "path", "p", "./", "Project path.")

	ansibleCmd := &cobra.Command{
		Use:   "ansible",
		Short: "ansible project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			o.Type = "ansible"
			return run(o)
		},
	}

	commandCmd := &cobra.Command{
		Use:   "mvc",
		Short: "mvc project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			o.Type = "mvc"
			return run(o)
		},
	}

	projectCmd.AddCommand(ansibleCmd)
	projectCmd.AddCommand(commandCmd)

	return cmd
}

func run(o *options.AppOptions) (err error) {
	server, err := o.NewServer()
	server.Run()
	return
}
