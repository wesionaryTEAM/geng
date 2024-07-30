package utils

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type FlagItem struct {
	Name  string
	Short string

	// Def default value passed in flag item, type is defaulted to string
	Def  any
	Desc string

	Persistent bool
}

// SetFlags sets the flags for the command
func SetFlags(cmd *cobra.Command, items []FlagItem) {
	for _, i := range items {

		if i.Def == nil {
			f := cmd.Flags()
			if i.Persistent {
				f = cmd.PersistentFlags()
			}

			f.StringP(i.Name, i.Short, "", i.Desc)
		}

		if def, ok := i.Def.([]string); ok {

			f := cmd.Flags()
			if i.Persistent {
				f = cmd.PersistentFlags()
			}

			f.StringSliceP(i.Name, i.Short, def, i.Desc)
		}

		if def, ok := i.Def.(string); ok {

			f := cmd.Flags()
			if i.Persistent {
				f = cmd.PersistentFlags()
			}

			f.StringP(i.Name, i.Short, def, i.Desc)
		}

	}
}

// BindFlag binds the flag to the viper
func BindFlag(cmd *cobra.Command, items []FlagItem) {
	for _, i := range items {
		f := cmd.Flags()
		if i.Persistent {
			f = cmd.PersistentFlags()
		}
		viper.BindPFlag(i.Name, f.Lookup(i.Name))
	}
}
