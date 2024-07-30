package pkg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BindStrFlag binds string flag to the passed command
func BindStrFlag(cmd *cobra.Command, key, short, def, desc string) {
	cmd.Flags().StringP(key, short, def, desc)
	viper.BindPFlag(key, cmd.Flags().Lookup(key))
}

// BindStrFlag binds string flag to the passed command
func BindStrPFlag(cmd *cobra.Command, key, short, def, desc string) {
	cmd.PersistentFlags().StringP(key, short, def, desc)
	viper.BindPFlag(key, cmd.PersistentFlags().Lookup(key))
}
