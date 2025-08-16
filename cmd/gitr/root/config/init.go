package config

import (
	"github.com/plantoncloud/gitr/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize gitr config",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := config.EnsureInitialConfig(); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success!")
}
