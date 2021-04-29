package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type input struct {
	cfgFile string
	dry     bool
}

var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "git rapid - the missing link b/w git cli & scm providers",
	Long:  `save time(a ton) by opening git repos on web browser right from the command line`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	i := &input{}
	cobra.OnInitialize(readConfig(i))
	rootCmd.PersistentFlags().StringVar(&i.cfgFile, "config", "", "config file (default is $HOME/.gitr.yaml)")
	rootCmd.PersistentFlags().BoolP("dry", "d", false, "dry run")
	rootCmd.PersistentFlags().BoolP("create-dir", "c", false, "create directories")
	viper.BindPFlag("dry", rootCmd.PersistentFlags().Lookup("dry"))
	viper.BindPFlag("create-dir", rootCmd.PersistentFlags().Lookup("create-dir"))
}

func readConfig(i *input) func() {
	return func() {
		if i.cfgFile != "" {
			viper.SetConfigFile(i.cfgFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			viper.AddConfigPath(home)
			viper.SetConfigName(".gitr")
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
		}
	}
}
