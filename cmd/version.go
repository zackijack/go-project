package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zackijack/go-project/internal/version"
)

const versionDesc = `
Show the version for go-project.
This will print a representation the version of go-project.
The output will look something like this:

Version    : v1.0.0
Git Commit : ff52399e51bb880526e9cd0ed8386f6433b74da1
Build Date : 2020-07-27-17:07:02
Go Version : go1.16.5
OS / Arch  : darwin / amd64

- Version is the semantic version of the release.
- Git Commit is the SHA for the commit that this version was built from.
- Build Date is the date when the application was built.
- Go Version is the version of the Go compiler used.
- OS / Arch is the version of the operating system and architecture was built for.
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Reveal go-project personal information",
	Long:  versionDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version    :", version.Version)
		fmt.Println("Git Commit :", version.GitCommit)
		fmt.Println("Build Date :", version.BuildDate)
		fmt.Println("Go Version :", version.GoVersion)
		fmt.Println("OS / Arch  :", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
