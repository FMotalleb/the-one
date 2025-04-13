/*
Copyright © 2025 Motalleb Fallahnezhad

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/fmotalleb/the-one/config"
	"github.com/fmotalleb/the-one/logging"
)

var (
	cfgFile string
	cfg     config.Config
	logCfg  logging.LogConfig
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "the-one",
	Short: "A simple init system for monolithic containers",
	Long: `Simple yet fast init system for monolithic containers.
It is designed to be lightweight and easy to use, making it ideal for
containers that require a simple init system.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := logging.BootLogger(logCfg); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initConfig(); err != nil {
			return err
		}
		data, err := json.Marshal(cfg)
		fmt.Printf("%s\n%v", data, err)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.seed.yaml)",
	)

	rootCmd.PersistentFlags().BoolVar(
		&logCfg.Development,
		"dev-logging",
		false,
		"enable verbose development logger instead of JSON",
	)

	rootCmd.PersistentFlags().BoolVar(
		&logCfg.ShowCaller,
		"	log-caller-info",
		false,
		"include caller filepath in log output",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	c, err := config.ReadAndMergeConfig(cfgFile)
	if err != nil {
		return err
	}
	cfg, err = config.DecodeConfig(c)
	if err != nil {
		return err
	}
	return nil
}
