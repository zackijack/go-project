package cmd

import (
	"github.com/zackijack/go-project/internal/helpers"
	"os"

	"github.com/spf13/cobra"
)

const completionDesc = `
To load completions:

Bash:

    $ source <(go-project completion bash)

    To load completions for each session, execute once:

    Linux:

      $ go-project completion bash > /etc/bash_completion.d/go-project

    MacOS:

      $ go-project completion bash > /usr/local/etc/bash_completion.d/go-project

Zsh:

    If shell completion is not already enabled in your environment you will need
    to enable it. You can execute the following once:

      $ echo "autoload -U compinit; compinit" >> ~/.zshrc

    To load completions for each session, execute once:

      $ go-project completion zsh > "${fpath[1]}/_go-project"

    You will need to start a new shell for this setup to take effect.

Fish:

    $ go-project completion fish | source

    To load completions for each session, execute once:

      $ go-project completion fish > ~/.config/fish/completions/go-project.fish

PowerShell:

    PS> go-project completion powershell | Out-String | Invoke-Expression

    To load completions for every new session, run:

      PS> go-project completion powershell > go-project.ps1

    And source this file from your PowerShell profile.
`

var errMsg = helpers.ErrMsg("completion generator")

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  completionDesc,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			helpers.CheckErr(cmd.Root().GenBashCompletion(os.Stdout), errMsg, true)
		case "zsh":
			helpers.CheckErr(cmd.Root().GenZshCompletion(os.Stdout), errMsg, true)
		case "fish":
			helpers.CheckErr(cmd.Root().GenFishCompletion(os.Stdout, true), errMsg, true)
		case "powershell":
			helpers.CheckErr(cmd.Root().GenPowerShellCompletion(os.Stdout), errMsg, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
