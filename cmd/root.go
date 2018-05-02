package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spacemeshos/go-spacemesh/api"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// SpacemeshApp is the cli app singleton
type SpacemeshApp struct {
	*cobra.Command
	Node             p2p.LocalNode
	NodeInitCallback chan bool
	grpcAPIService   *api.SpaceMeshGrpcService
	jsonAPIService   *api.JSONHTTPServer
}

func before(cmd *cobra.Command, args []string) { fmt.Printf("Called before\n") }

func cleanup(cmd *cobra.Command, args []string) { fmt.Printf("Called cleanup\n") }

func startSpacemeshNode(cmd *cobra.Command, args []string) { fmt.Printf("Called startSpacemeshNode\n") }

var rootCmd = &cobra.Command{
	Use:   "newApp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun:  before,
	Run:     cleanup,
	PostRun: startSpacemeshNode,
}

// App is main app entry point.
// It provides access the local node and other top-level modules.
var App = &SpacemeshApp{rootCmd, nil, make(chan bool, 1), nil, nil}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := App.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	App.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	App.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".newApp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
