package tunnel

import (
	"github.com/spf13/cobra"

	"github.com/fatih/color"
)

func rootMain(command *cobra.Command, arguments []string) error {
	// If no commands were given, then print help information and bail. We don't
	// have to worry about warning about arguments being present here (which
	// would be incorrect usage) because arguments can't even reach this point
	// (they will be mistaken for subcommands and a error will be displayed).
	command.Help()

	// Success.
	return nil
}

var RootCommand = &cobra.Command{
	Use:          "tunnel",
	Short:        "Create and manage tunnels",
	RunE:         rootMain,
	SilenceUsage: true,
}

var rootConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Mark the command as experimental.
	RootCommand.Short = RootCommand.Short + color.YellowString(" [Experimental]")

	// Grab a handle for the command line flags.
	flags := RootCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&rootConfiguration.help, "help", "h", false, "Show help information")

	// Register commands.
	RootCommand.AddCommand(
		createCommand,
		listCommand,
		monitorCommand,
		pauseCommand,
		resumeCommand,
		terminateCommand,
		hostCommand,
	)
}
