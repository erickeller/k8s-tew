package main

import (
	"os"

	"github.com/darxkies/k8s-tew/servers"
	"github.com/darxkies/k8s-tew/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var killTimeout uint

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run",
	Long:  "Run servers",
	Run: func(cmd *cobra.Command, args []string) {
		if error := Bootstrap(true); error != nil {
			log.WithFields(log.Fields{"error": error}).Error("Failed initialization")

			os.Exit(-1)
		}

		if len(_config.Config.Nodes) == 0 {
			log.WithFields(log.Fields{"error": "no nodes defined"}).Error("Failed to run")

			os.Exit(-1)
		}

		if _config.Node == nil {
			log.WithFields(log.Fields{"error": "current host not found in the list of nodes"}).Error("Failed to run")

			os.Exit(-1)
		}

		serversContainer := servers.NewServers(_config, killTimeout)

		utils.SetProgressSteps(serversContainer.Steps())

		utils.ShowProgress()

		if error := serversContainer.Run(commandRetries); error != nil {
			log.WithFields(log.Fields{"error": error}).Error("Failed to run")

			os.Exit(-1)
		}
	},
}

func init() {
	runCmd.Flags().UintVarP(&commandRetries, "command-retries", "r", 300, "The count of command retries")
	runCmd.Flags().UintVar(&killTimeout, "kill-timeout", 10, "Kill timeout for child processes")
	RootCmd.AddCommand(runCmd)
}
