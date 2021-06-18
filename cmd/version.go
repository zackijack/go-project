package cmd

import (
	"fmt"

	"github.com/zackijack/go-project/internal/version"
	"github.com/spf13/cobra"
)

const versionDesc = `
Showing who exactly go-project is.
This will print a representation profile of go-project.
The output will look something like this:

Version: v1.0.0
Git Commit: ff52399e51bb880526e9cd0ed8386f6433b74da1
Build Date: 2020-07-27-17:07:02
Go Version: go1.14.6
OS Arch: darwin amd64

- Version is the semantic version of the release.
- Git Commit is the SHA for the commit that this version was built from.
- Build Date is the date when the application was built.
- Go Version is the version of the Go compiler used.
- OS Arch is the version of the operating system and architecture was built for.
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Reveal go-project personal information",
	Long:  versionDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS Arch:", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
