package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

    "github.com/kennethdyer/versc/logger"
    "github.com/kennethdyer/versc/search"
)

const ver string = "0.1.0"

var (

	// CommitStamp - Provides shortened form of last commit on repository at build-time.
	CommitStamp string

	// BuildAt - Provides timestamp of last build.
	BuildAt string
)

/*
reportVersion function takes a boolean argument to indicate verobsity, and then
prints to stdout information on the current version and built of Avocet Tools.
*/
func reportVersion(verbose bool) {
	var content []string
	if verbose {
		content = []string{
			"Versc - Version Scanner",
			"Kenneth P. J. Dyer <kenneth.dyer@mongodb.com>",
			fmt.Sprintf("Version: v%s", ver),
			fmt.Sprintf("Commit: %s", CommitStamp),
			fmt.Sprintf("Build Time: %s", BuildAt),
		}
	} else {
		content = []string{fmt.Sprintf("versc - version %s\n", ver)}
	}
	fmt.Println(strings.Join(content, "\n  "))
	os.Exit(0)
}

func init() {
	// Logging Defaults
	viper.SetDefault("verbosity", false)
	viper.SetDefault("debug", false)
	viper.SetDefault("trace", false)

	// Parallel Defaults
	viper.SetDefault("parallel", runtime.NumCPU())

    // Include Patterns
    viper.SetDefault("exclude_pattern", "release-notes")
    viper.SetDefault("include_extensions", ".rst,.txt,.yml,.yaml")


	// Set and Read Configuration File
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/avocet")

	viper.SetConfigName("versc.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}

}

// Initialize Commands
var cmd = &cobra.Command{
	Use:     "versc",
	Short:   "Scans documentation projects for version descriptions",
	Version: ver,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 1 {
            search.RunSearch(args[0], []string{})
        } else {
            search.RunSearch(args[0], args[1:])
        }
    },
}


func main() {
	/******************************** OPTIONS *************************************/
	// Verbosity Flag
	cmd.PersistentFlags().BoolP(
		"verbose", "v",
		viper.GetBool("verbosity"),
		"Enables verbose output")
	viper.BindPFlag("verbosity", cmd.PersistentFlags().Lookup("verbose"))

	// Debug
	cmd.PersistentFlags().BoolP(
		"debug", "D",
		viper.GetBool("debug"),
		"Enables debugging output for logs")
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	// Trace
	cmd.PersistentFlags().BoolP(
		"trace", "T",
		viper.GetBool("trace"),
		"Enables trace output for logs")
	viper.BindPFlag("trace", cmd.PersistentFlags().Lookup("trace"))

	// Parallel Processing
	cmd.PersistentFlags().IntP(
		"jobs", "j",
		viper.GetInt("parallel"),
		"Specifies the number of threads to utilize for specific tasks")
	viper.BindPFlag("parallel", cmd.PersistentFlags().Lookup("jobs"))

	cmd.Execute()

}
