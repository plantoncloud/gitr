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
	debug   bool
}

var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "git rapid - the missing link b/w git cli & scm providers",
	Long: `tool to navigate to important features of scm efficiently right from the command line.
no more searching for clone url, simply use the browser url to clone repos.`,
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
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "verbose logging")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
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
