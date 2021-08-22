package gitr

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"gopkg.in/yaml.v3"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "see gitr config",
	Run:   configHandler,
}

func configHandler(cmd *cobra.Command, args []string) {
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
