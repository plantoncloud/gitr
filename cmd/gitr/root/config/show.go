package config

import (
	"fmt"
	"github.com/plantoncloud/gitr/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var Show = &cobra.Command{
	Use:   "show",
	Short: "show gitr config",
	Run:   showHandler,
}

func showHandler(cmd *cobra.Command, args []string) {
	cfg, err := config.NewGitrConfig()
	if err != nil {
		log.Fatalf("failed to get gitr config. err: %v", err)
	}
	d, err := yaml.Marshal(&cfg)
	fmt.Printf("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", string(d))
}
