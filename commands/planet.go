package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
var RootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: `Dagobah is an awesome planet style RSS aggregator`,
	Long:  `RSS agg + YAML and web blah blah`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dagobah runs")
	},
}


func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
*/

var CfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/dagobah/")
	viper.AddConfigPath("$HOME/.dagobah/")
	viper.ReadInConfig()
}

var RootCmd = &cobra.Command{
	Use:   "",
	Short: `...`,
	Long:  `...`,
	Run:   rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println(viper.Get("feeds"))
	fmt.Println(viper.GetString("appname"))
}

func addCommands() {
	RootCmd.AddCommand(fetchCmd)
}

func Execute() {
	addCommands()

	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
