package config

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/leftbin/go-util/pkg/shell"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "edit gitr config",
	Run:   editHandler,
}

func editHandler(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get home dir. err: %v", err)
	}
	gitrConfigPath := filepath.Join(homeDir, ".gitr.yaml")
	if !file.IsFileExists(gitrConfigPath) {
		log.Fatalf("config file %s not found. run 'gitr config init' to create one", gitrConfigPath)
	}
	if err := shell.RunCmd(exec.Command("code", gitrConfigPath)); err != nil {
		log.Fatalf("failed to open gitr config. err: %v", err)
	}
}
