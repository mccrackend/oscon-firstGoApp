package commands

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Feeds []string
	Port  int
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch feeds",
	Long:  `Dagobah will fetch all feeds listed in the config file.`,
	Run:   fetchRun,
}

func init() {
	fetchCmd.Flags().Int("rsstimeout", 5, "Timeout (in min) for RSS retrival")
	viper.BindPFlag("rsstimeout", fetchCmd.Flags().Lookup("rsstimeout"))
}

func fetchRun(cmd *cobra.Command, args []string) {
	Fetcher()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

func Fetcher() {
	var config Config

	if err := viper.Marshal(&config); err != nil {
		fmt.Println(err)
	}

	for _, feed := range config.Feeds {
		go PollFeed(feed)
	}
}

func PollFeed(string feed) {

}
