package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var (
	version string = "0.0.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of yyds",
	Long:  "All soft has versions ,this is the yyds's,print the version for you",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version())
	},
}

func Version() string {
	command := exec.Command("go", "version")
	output, err := command.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	return version
}
