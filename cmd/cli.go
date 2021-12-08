package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/danstn/moped/cmd/register"
	"github.com/danstn/moped/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type MopedCLI struct {
	rootCmd *cobra.Command
}

func NewMopedCLI() *MopedCLI {
	mopedCLI := &MopedCLI{}

	// init config
	cobra.OnInitialize(initConfig)

	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("failed loading app config: %v", err)
	}

	// root command
	mopedCLI.rootCmd = initRootCmd(appConfig)
	return mopedCLI
}

func initRootCmd(appConfig *config.AppConfig) *cobra.Command {
	// root command
	rootCmd := &cobra.Command{
		Use:   "moped",
		Short: "Moped - pocket CI.",
		Long:  `Moped is your pocket CI. Because you can.`,
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.moped.yaml)")

	// register command
	registerCmd := register.NewCommand(appConfig)

	// attach commands
	rootCmd.AddCommand(registerCmd)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (cli *MopedCLI) Execute() {
	cobra.CheckErr(cli.rootCmd.Execute())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".moped" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".moped")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
